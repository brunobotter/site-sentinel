package service

import (
	"context"
	"time"

	"github.com/brunobotter/site-sentinel/application/domain"
	apphttp "github.com/brunobotter/site-sentinel/application/http"
	"github.com/brunobotter/site-sentinel/application/repo"
	"github.com/brunobotter/site-sentinel/application/service"
)

type checkExecutionService struct {
	planner    service.MonitorPlannerService
	httpClient apphttp.Client
	resultRepo repo.CheckResultRepository
}

func NewCheckExecutionService(
	planner service.MonitorPlannerService,
	httpClient apphttp.Client,
	resultRepo repo.CheckResultRepository,
) *checkExecutionService {
	return &checkExecutionService{
		planner:    planner,
		httpClient: httpClient,
		resultRepo: resultRepo,
	}
}

func (s *checkExecutionService) RunBatch(ctx context.Context, targets []domain.MonitorTarget) error {
	batches := s.planner.PlanBatch(targets)
	for _, batch := range batches {
		results := make([]domain.CheckResult, 0, len(batch))
		for _, target := range batch {
			results = append(results, s.checkTarget(ctx, target))
		}
		if err := s.resultRepo.SaveBatch(ctx, results); err != nil {
			return err
		}
	}

	return nil
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
