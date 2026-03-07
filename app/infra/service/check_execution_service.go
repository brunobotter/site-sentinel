package service

import (
	"context"
	"sync"
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
	apphttp "github.com/brunobotter/site-sentinel/application/http"
	"github.com/brunobotter/site-sentinel/application/repo"
	"github.com/brunobotter/site-sentinel/application/service"
)

const (
	defaultWorkerCount = 50
	defaultQueueSize   = 200
)

type CheckExecutionSettings struct {
	WorkerCount int
	QueueSize   int
}
type checkExecutionService struct {
	planner    service.MonitorPlannerService
	httpClient apphttp.Client
	resultRepo repo.CheckResultRepository
	settings   CheckExecutionSettings
}

func NewCheckExecutionService(
	planner service.MonitorPlannerService,
	httpClient apphttp.Client,
	resultRepo repo.CheckResultRepository,
	settings CheckExecutionSettings,
) *checkExecutionService {
	if settings.WorkerCount <= 0 {
		settings.WorkerCount = defaultWorkerCount
	}
	if settings.QueueSize <= 0 {
		settings.QueueSize = defaultQueueSize
	}

	return &checkExecutionService{
		planner:    planner,
		httpClient: httpClient,
		resultRepo: resultRepo,
		settings:   settings,
	}
}

func (s *checkExecutionService) RunBatch(ctx context.Context, targets []domain.MonitorTarget) error {
	batches := s.planner.PlanBatch(targets)
	for _, batch := range batches {
		results, err := s.runConcurrentChecks(ctx, batch)
		if err != nil {
			return err
		}
		if err := s.resultRepo.SaveBatch(ctx, results); err != nil {
			return err
		}
	}

	return nil
}
func (s *checkExecutionService) runConcurrentChecks(ctx context.Context, targets []domain.MonitorTarget) ([]domain.CheckResult, error) {
	if len(targets) == 0 {
		return []domain.CheckResult{}, nil
	}

	jobs := make(chan domain.MonitorTarget, s.settings.QueueSize)
	results := make(chan domain.CheckResult, len(targets))

	workerCount := s.settings.WorkerCount
	if workerCount > len(targets) {
		workerCount = len(targets)
	}

	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case target, ok := <-jobs:
					if !ok {
						return
					}
					results <- s.checkTarget(ctx, target)
				}
			}
		}()
	}

	for _, target := range targets {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return nil, ctx.Err()
		case jobs <- target:
		}
	}
	close(jobs)
	wg.Wait()
	close(results)

	checks := make([]domain.CheckResult, 0, len(targets))
	for result := range results {
		checks = append(checks, result)
	}

	return checks, nil
}
func (s *checkExecutionService) checkTarget(ctx context.Context, target domain.MonitorTarget) domain.CheckResult {
	method := target.Method
	if method == "" {
		method = "GET"
	}

	timeout := target.Policy.Timeout
	if timeout <= 0 {
		timeout = 3 * time.Second
	}

	retries := target.Policy.Retries
	if retries < 0 {
		retries = 0
	}

	var (
		statusCode   int
		responseTime time.Duration
		isUp         bool
		lastErr      string
	)

	for attempt := 0; attempt <= retries; attempt++ {
		attemptCtx, cancel := context.WithTimeout(ctx, timeout)
		startedAt := time.Now()

		req, err := s.httpClient.NewRequestWithContext(attemptCtx, method, target.URL, nil)
		if err != nil {
			responseTime = time.Since(startedAt)
			lastErr = err.Error()
			cancel()
			continue
		}

		resp, err := s.httpClient.Do(attemptCtx, "monitor-check", req)
		responseTime = time.Since(startedAt)
		cancel()
		if err != nil {
			lastErr = err.Error()
			if attempt < retries && target.Policy.RetryDelay > 0 {
				time.Sleep(target.Policy.RetryDelay)
			}
			continue
		}

		statusCode = resp.Status()
		isUp = statusCode == target.Policy.ExpectedStatus
		lastErr = ""
		break
	}

	result := domain.NewCheckResult(target.ID, statusCode, responseTime, isUp, lastErr)
	return *result
}
