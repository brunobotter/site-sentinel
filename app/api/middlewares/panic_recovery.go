package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/labstack/echo/v4"
)

type PanicMiddleware struct {
	middlewareFunc any
}

func (m *PanicMiddleware) GetMiddleware() any {
	return m.middlewareFunc
}

func NewPanicMiddleware(logger logger.Logger) MiddlewareFunc {
	return &PanicMiddleware{
		middlewareFunc: getPanicRecoveryMiddlewareFunc(logger),
	}
}

func getPanicRecoveryMiddlewareFunc(log logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func() {
				if rec := recover(); rec != nil {
					stackTrace := string(debug.Stack())
					var panicMessage string
					switch v := rec.(type) {
					case string:
						panicMessage = v
					case error:
						panicMessage = v.Error()
					default:
						panicMessage = "unknown panic"
					}

					log := logger.LoggerFromContext(ctx.Request().Context())

					log.Error(fmt.Errorf("PANIC RECOVERED IN REQUEST - Path: %s, Method: %s, Error: %v",
						ctx.Request().URL.Path,
						ctx.Request().Method,
						panicMessage))

					log.Error(fmt.Errorf("stack trace: %s", stackTrace))

					ctx.Error(echo.NewHTTPError(http.StatusInternalServerError, "tente novamente mais tarde."))
				}
			}()

			return next(ctx)
		}
	}
}
