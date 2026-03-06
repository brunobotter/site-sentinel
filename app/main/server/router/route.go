package router

import (
	"github.com/brunobotter/site-sentinel/api/middlewares"
)

type Router interface {
	GET(path string, controller any, m ...middlewares.MiddlewareFunc)
	POST(path string, controller any, m ...middlewares.MiddlewareFunc)
	PUT(path string, controller any, m ...middlewares.MiddlewareFunc)
	PATCH(path string, controller any, m ...middlewares.MiddlewareFunc)
	DELETE(path string, controller any, m ...middlewares.MiddlewareFunc)
	Use(m ...middlewares.MiddlewareFunc)
	Group(prefix string, groupingFn func(group RouteGroup), m ...middlewares.MiddlewareFunc) RouteGroup
}

type RouteGroup interface {
	Router
}
