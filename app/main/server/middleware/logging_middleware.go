package middleware

import (
	"time"

	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/labstack/echo/v4"
)

func RequestLogger(log logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()

			err := next(c)
			duration := time.Since(start)
			status := c.Response().Status
			if err != nil {
				log.Errorf(
					"request_error method=%s path=%s query=%s status=%d duration_ms=%d remote_ip=%s user_agent=%s error=%v",
					req.Method,
					c.Path(),
					req.URL.RawQuery,
					status,
					duration.Milliseconds(),
					c.RealIP(),
					req.UserAgent(),
					err,
				)
				return err
			}

			if status >= 400 {
				log.Errorf(
					"request_failed method=%s path=%s query=%s status=%d duration_ms=%d remote_ip=%s user_agent=%s",
					req.Method,
					c.Path(),
					req.URL.RawQuery,
					status,
					duration.Milliseconds(),
					c.RealIP(),
					req.UserAgent(),
				)
			}
			return nil
		}
	}
}
