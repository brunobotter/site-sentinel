package providers

import (
	"github.com/brunobotter/site-sentinel/application/repo"
	"github.com/brunobotter/site-sentinel/infra/logger"
	infraRepo "github.com/brunobotter/site-sentinel/infra/repo"
	"github.com/brunobotter/site-sentinel/main/container"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryProvider struct{}

func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{}
}
func (p *RepositoryProvider) Register(c container.Container) {
	c.Singleton(func(db *pgxpool.Pool, log logger.Logger) repo.MonitorTargetRepository {
		return infraRepo.NewMonitorTargetPgRepository(db, log)
	})

	c.Singleton(func(db *pgxpool.Pool, log logger.Logger) repo.CheckResultRepository {
		return infraRepo.NewCheckResultPgRepository(db, log)
	})
}
