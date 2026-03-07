package service

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/brunobotter/site-sentinel/application/repo"
	appservice "github.com/brunobotter/site-sentinel/application/service"
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/jackc/pgx/v5/pgconn"
)

const defaultSchedulerInterval = 60 * time.Second

type monitorSchedulerService struct {
	targetRepo     repo.MonitorTargetRepository
	checkExecution appservice.CheckExecutionService
	log            logger.Logger
	interval       time.Duration
	enabled        bool
	isCycleRunning atomic.Bool
	schemaWarned   atomic.Bool
}

func NewMonitorSchedulerService(
	targetRepo repo.MonitorTargetRepository,
	checkExecution appservice.CheckExecutionService,
	log logger.Logger,
	interval time.Duration,
	enabled bool,
) *monitorSchedulerService {
	if interval <= 0 {
		interval = defaultSchedulerInterval
	}

	return &monitorSchedulerService{
		targetRepo:     targetRepo,
		checkExecution: checkExecution,
		log:            log,
		interval:       interval,
		enabled:        enabled,
	}
}

func (s *monitorSchedulerService) Start(ctx context.Context) {
	if !s.enabled {
		s.log.Infof("monitor scheduler desabilitado")
		return
	}

	s.log.Infof("monitor scheduler iniciado (intervalo=%s)", s.interval)
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	s.runCycle(ctx)

	for {
		select {
		case <-ctx.Done():
			s.log.Infof("monitor scheduler finalizado")
			return
		case <-ticker.C:
			s.runCycle(ctx)
		}
	}
}

func (s *monitorSchedulerService) runCycle(ctx context.Context) {
	if !s.isCycleRunning.CompareAndSwap(false, true) {
		s.log.Debugf("monitor scheduler: ciclo anterior ainda executando, pulando")
		return
	}
	defer s.isCycleRunning.Store(false)

	targets, err := s.targetRepo.ListActive(ctx)
	if err != nil {
		if isUndefinedTableError(err) {
			s.logMissingSchema()
			return
		}

		s.log.Errorf("monitor scheduler: erro ao listar targets ativos: %v", err)
		return
	}
	s.schemaWarned.Store(false)
	if len(targets) == 0 {
		s.log.Debugf("monitor scheduler: nenhum target ativo")
		return
	}

	if err := s.checkExecution.RunBatch(ctx, targets); err != nil {
		s.log.Errorf("monitor scheduler: erro ao executar batch: %v", err)
		return
	}

	s.log.Infof("monitor scheduler: ciclo concluído com %d targets", len(targets))
}

func (s *monitorSchedulerService) logMissingSchema() {
	if !s.schemaWarned.CompareAndSwap(false, true) {
		s.log.Debugf("monitor scheduler: aguardando migrations para monitor_targets")
		return
	}

	s.log.Errorf("monitor scheduler: tabela monitor_targets nao existe; execute as migrations em /migrations antes de iniciar o monitoramento")
}

func isUndefinedTableError(err error) bool {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}

	return pgErr.Code == "42P01"
}
