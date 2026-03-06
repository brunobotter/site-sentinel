package adapters

import (
	"bytes"
	"io"
	"reflect"

	_http "net/http"

	"github.com/brunobotter/site-sentinel/api/http"
	"github.com/brunobotter/site-sentinel/api/middlewares"
	"github.com/brunobotter/site-sentinel/main/server/router"
	"github.com/labstack/echo/v4"
)

// TODO: Adicionar tracking (observabilidade) para subir o erro_trace
type echoRouterAdapter struct {
	echo *echo.Group
}

func NewEchoRouterAdapter(e *echo.Echo) router.Router {
	return &echoRouterAdapter{
		echo: e.Group(""),
	}
}

func (a *echoRouterAdapter) GET(path string, handler any, m ...middlewares.MiddlewareFunc) {
	a.echo.GET(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) POST(path string, handler any, m ...middlewares.MiddlewareFunc) {
	a.echo.POST(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) PUT(path string, handler any, m ...middlewares.MiddlewareFunc) {
	a.echo.PUT(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) PATCH(path string, handler any, m ...middlewares.MiddlewareFunc) {
	a.echo.PATCH(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) DELETE(path string, handler any, m ...middlewares.MiddlewareFunc) {
	a.echo.DELETE(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) Use(m ...middlewares.MiddlewareFunc) {
	a.echo.Use(a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) Group(prefix string, groupingFn func(group router.RouteGroup), m ...middlewares.MiddlewareFunc) router.RouteGroup {
	echoGroup := a.echo.Group(prefix, a.adaptEchoMiddlewares(m)...)
	groupRouter := &echoRouterAdapter{
		echo: echoGroup,
	}

	if groupingFn != nil {
		groupingFn(groupRouter)
	}

	return groupRouter
}

func (a *echoRouterAdapter) adaptPath(path string) string {
	if path == "/" {
		return ""
	}
	return path
}

func (a *echoRouterAdapter) adaptEchoRoute(handler any) func(c echo.Context) error {
	refHandler := reflect.ValueOf(handler)

	return func(c echo.Context) error {
		request := a.buildRequest(c, &refHandler)

		handlerArgs := []reflect.Value{request}
		handlerResponse := refHandler.Call(handlerArgs)

		if len(handlerResponse) == 0 {
			return c.JSON(_http.StatusNoContent, nil)
		}

		response := handlerResponse[0].Interface().(*http.HttpResponse)
		// if response.ErrorMessage != "" {
		// 	c.Set(application.ErrorMessageContextKey, response.ErrorMessage)
		// }

		//if response.ErrorMessage != "" {
		//	if a.tracking != nil {
		//		a.tracking.CreateAttribute(c.Request().Context(), "error_trace", response.ErrorMessage)
		//	}
		//}

		return c.JSON(response.StatusCode, response.Body)
	}
}

func (a *echoRouterAdapter) buildRequest(c echo.Context, refHandler *reflect.Value) reflect.Value {
	baseRequest := a.adaptEchoRequest(c)
	request := reflect.ValueOf(baseRequest)

	if refHandler.Type().NumIn() > 0 {
		handlerArg := refHandler.Type().In(0)

		if handlerArg.Name() == "HttpRequest" {
			return request
		}

		println(handlerArg.Name())

		requestArgInstance := reflect.Zero(handlerArg.Elem()).Interface()
		v := reflect.ValueOf(&requestArgInstance).Elem()
		request = reflect.New(v.Elem().Type()).Elem()

		field := request.FieldByName("HttpRequest")

		if !field.IsValid() {
			field = request.FieldByName("AuthenticatedRequest")
		}

		if !field.IsValid() {
			panic("request does not correct implement HttpRequest or AuthenticatedRequest")
		}

		field.Set(reflect.ValueOf(baseRequest))

		reqBody, _ := io.ReadAll(c.Request().Body)

		c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody))

		c.Bind(request.Addr().Interface())

		c.Request().Body = io.NopCloser(bytes.NewBuffer(reqBody))

	}

	return request.Addr()
}
func (a *echoRequestAdapter) UserAgent() string {
	return a.echo.Request().UserAgent()
}
func (a *echoRouterAdapter) adaptEchoRequest(context echo.Context) http.HttpRequest {
	return NewEchoRequestAdapter(context)
}

func (a *echoRouterAdapter) adaptEchoMiddlewares(handlers []middlewares.MiddlewareFunc) []echo.MiddlewareFunc {
	m := []echo.MiddlewareFunc{}
	for _, handler := range handlers {
		m = append(m, handler.GetMiddleware().(echo.MiddlewareFunc))
	}
	return m
}
