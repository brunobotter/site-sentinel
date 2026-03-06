package middlewares

import (
	"net/http"
	"strings"

	"github.com/brunobotter/site-sentinel/main/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSMiddleware struct {
	middlewareFunc any
}

func (m *CORSMiddleware) GetMiddleware() any {
	return m.middlewareFunc
}

func NewCORSMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	/*origins := parseOrigins(cfg.FrontURLs)
	if len(origins) == 0 {
		return blockAllCORS()
	}*/
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions},
		AllowHeaders:     []string{"content-type", "x-itau-apikey", "x-itau-correlationid", "x-charon-params", "x-charon-session", "x-itau-recaptcha-token", "x-journey-token", "x-journey-id", "user-agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           600,
	})
}

func parseOrigins(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	fields := strings.Fields(raw)
	out := make([]string, 0, len(fields))
	for _, f := range fields {
		o := strings.TrimSpace(f)
		if o != "" {
			out = append(out, o)
		}
	}
	return out
}

func blockAllCORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			if origin != "" {
				return c.NoContent(http.StatusForbidden)
			}
			return next(c)
		}
	}
}
