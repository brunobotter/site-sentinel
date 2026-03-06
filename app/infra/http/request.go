package http

import (
	"bytes"
	"context"
	"io"
	_http "net/http"

	httpContract "github.com/brunobotter/site-sentinel/application/http"
)

type request struct {
	Req  *_http.Request
	Body []byte
}

func newRequest(method string, url string, body []byte) (httpContract.Request, error) {
	var payload io.Reader = nil
	if body != nil {
		payload = bytes.NewBuffer(body)
	}
	req, err := _http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	return &request{req, body}, nil
}

func newRequestWithContext(ctx context.Context, method string, url string, body []byte) (httpContract.Request, error) {
	var payload io.Reader = nil
	if body != nil {
		payload = bytes.NewReader(body)
	}
	req, err := _http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return nil, err
	}
	return &request{req, body}, nil
}

func (r *request) SetHeader(key, value string) {
	r.Req.Header.Set(key, value)
}

func (r *request) WithContext(ctx context.Context) httpContract.Request {
	r.Req = r.Req.WithContext(ctx)
	return r
}
