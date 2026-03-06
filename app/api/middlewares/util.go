package middlewares

import (
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/brunobotter/site-sentinel/main/config"
)

func CommonMiddlewares(logger logger.Logger, cfg *config.Config) []MiddlewareFunc {
	return []MiddlewareFunc{
		NewPanicMiddleware(logger),
		NewLoggerMiddleware(logger, cfg),
	}
}
