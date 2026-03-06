package middlewares

import (
	"strings"

	l "github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/brunobotter/site-sentinel/main/config"
	"github.com/brunobotter/site-sentinel/util/shared"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	ua "github.com/mileusna/useragent"
)

const (
	CorrelationID string = "x-itau-correlationID"
)

type LoggerMiddleware struct {
	middlewareFunc any
}

func (m *LoggerMiddleware) GetMiddleware() any {
	return m.middlewareFunc
}

func NewLoggerMiddleware(logger l.Logger, config *config.Config) MiddlewareFunc {
	return &LoggerMiddleware{
		middlewareFunc: getLoggerMiddlewareFunc(logger, config),
	}
}

func getLoggerMiddlewareFunc(logger l.Logger, config *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			userAgent := ctx.Request().UserAgent()

			if userAgent == "" {
				userAgent = ctx.Request().Header.Get("user-agent")

			}

			isMobile := IsFromMobile(ctx, userAgent)
			correlationID := strings.TrimSpace(ctx.Request().Header.Get(CorrelationID))

			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			log := logger.WithContext(ctx.Request().Context()).WithFields(map[string]any{
				"path":           ctx.Path(),
				"remote_ip":      ctx.RealIP(),
				"is_mobile":      isMobile,
				"correlation_id": correlationID,
			})

			appContext := ctx.Request().Context()
			appContext = l.SetContextLogger(appContext, log)
			appContext = shared.SetUserAgent(appContext, userAgent)
			appContext = shared.SetIP(appContext, ctx.RealIP())
			appContext = shared.SetIsFromMobile(appContext, isMobile)
			appContext = shared.SetCorrelationID(appContext, correlationID)
			appContext = shared.SetContextApplicationName(appContext, config.App_Name)
			appContext = shared.SetContextApplicationEnvironment(appContext, config.Env)
			ctx.SetRequest(ctx.Request().WithContext(appContext))
			return next(ctx)
		}
	}
}

func IsFromMobile(c echo.Context, userAgent string) bool {
	ua := ua.Parse(userAgent)
	return ua.Mobile || ua.Tablet

}
