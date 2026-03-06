package middlewares

type MiddlewareFunc interface {
	GetMiddleware() any
}
