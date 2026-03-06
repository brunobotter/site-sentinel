package adapters

import (
	"bytes"
	"context"
	"io"

	"github.com/labstack/echo/v4"
)

type echoRequestAdapter struct {
	echo echo.Context
}

func NewEchoRequestAdapter(echo echo.Context) *echoRequestAdapter {
	return &echoRequestAdapter{
		echo: echo,
	}
}

func (a *echoRequestAdapter) Context() context.Context {
	return a.echo.Request().Context()
}

func (a *echoRequestAdapter) Bind(i interface{}) error {
	return a.echo.Bind(i)
}

func (a *echoRequestAdapter) Param(name string) string {
	return a.echo.Param(name)
}

func (a *echoRequestAdapter) QueryParam(name string) string {
	return a.echo.QueryParam(name)
}

func (a *echoRequestAdapter) GetHeader(name string) string {
	return a.echo.Request().Header.Get(name)
}

func (a *echoRequestAdapter) Body() []byte {
	b := a.echo.Request().Body
	body, _ := io.ReadAll(b)
	a.echo.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	return body
}

func (a *echoRequestAdapter) Method() string {
	return a.echo.Request().Method
}

func (a *echoRequestAdapter) Path() string {
	return a.echo.Request().URL.Path
}
