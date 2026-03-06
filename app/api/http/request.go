package http

import (
	"context"
)

type HttpRequest interface {
	Context() context.Context
	Bind(i interface{}) error
	Param(name string) string
	QueryParam(name string) string
	GetHeader(name string) string
	Body() []byte
	Method() string
	Path() string
	UserAgent() string
}
