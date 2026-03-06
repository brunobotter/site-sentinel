package usecase

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/command"
	"github.com/brunobotter/site-sentinel/application/domain"
)

type CreateTargetUseCase interface {
	Execute(ctx context.Context, cmd command.CreateTargetCommand) error
}
type ListTargetsUseCase interface {
	Execute(ctx context.Context) ([]domain.MonitorTarget, error)
}
type RunBatchCheckUseCase interface {
	Execute(ctx context.Context, cmd command.RunCheckBatchCommand) error
}
type ListLatestResultsUseCase interface {
	Execute(ctx context.Context, limit int) ([]domain.CheckResult, error)
}
