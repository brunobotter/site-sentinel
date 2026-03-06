package providers

import (
	"context"
	"fmt"

	"github.com/brunobotter/site-sentinel/main/config"
	"github.com/brunobotter/site-sentinel/main/container"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseServiceProvider struct{}

func NewDatabaseServiceProvider() *DatabaseServiceProvider {
	return &DatabaseServiceProvider{}
}

func (p *DatabaseServiceProvider) Register(c container.Container) {
	c.Singleton(func(cfg *config.Config) *pgxpool.Pool {
		pool, err := pgxpool.New(context.Background(), cfg.Database.URL)
		if err != nil {
			panic(fmt.Errorf("nao pode inicializar conexao com banco: %w", err))
		}
		return pool
	})
}

func (p *DatabaseServiceProvider) Shutdown(pool *pgxpool.Pool) {
	pool.Close()
}
