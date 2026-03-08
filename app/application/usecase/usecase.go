package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/domain"
)

type CreateTargetUseCase interface {
	// Execute recebe um comando e cria um novo alvo de monitoramento.
	Execute(ctx context.Context, cmd command.CreateTargetCommand) error
}
type ListTargetsUseCase interface {
	// Execute busca os alvos cadastrados para exibição e processamento.
	Execute(ctx context.Context) ([]domain.MonitorTarget, error)
}
type RunBatchCheckUseCase interface {
	// Execute dispara os checks em lote para uma lista de alvos.
	Execute(ctx context.Context, cmd command.RunCheckBatchCommand) error
}
type ListLatestResultsUseCase interface {
	// Execute retorna os últimos resultados globais com limite configurável.
	Execute(ctx context.Context, limit int) ([]domain.CheckResult, error)
}
