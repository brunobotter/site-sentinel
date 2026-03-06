package repo

import (
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantPgRepo struct {
	db  *pgxpool.Pool
	log logger.Logger
}
