package repo

import (
	"context"

	"github.com/brunobotter/site-sentinel/application/domain"
)

type MonitorTargetRepository interface {
	Create(ctx context.Context, target domain.MonitorTarget) error
	Update(ctx context.Context, target domain.MonitorTarget) error
	Delete(ctx context.Context, id string) error

	FindByID(ctx context.Context, id string) (*domain.MonitorTarget, error)

	List(ctx context.Context, limit int, offset int) ([]domain.MonitorTarget, error)

	ListActive(ctx context.Context) ([]domain.MonitorTarget, error)
}

type CheckResultRepository interface {
	Save(ctx context.Context, result domain.CheckResult) error

	SaveBatch(ctx context.Context, results []domain.CheckResult) error

	ListLatestByTarget(
		ctx context.Context,
		targetID string,
		limit int,
	) ([]domain.CheckResult, error)

	ListLatestGlobal(
		ctx context.Context,
		limit int,
	) ([]domain.CheckResult, error)
}
