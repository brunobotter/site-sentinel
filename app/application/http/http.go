package http

import (
	"context"
)

type Client interface {
	NewRequestWithContext(ctx context.Context, method string, url string, body []byte) (request Request, err error)
	NewRequest(method string, url string, body []byte) (request Request, err error)
	Do(ctx context.Context, service string, req Request) (response Response, err error)
}

type Request interface {
	SetHeader(key, value string)
	WithContext(ctx context.Context) Request
}

type Response interface {
	Status() (status int)
	Body() (body []byte, err error)
	Header(key string) map[string][]string
}
