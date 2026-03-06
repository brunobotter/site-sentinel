package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	_http "net/http"
	"net/url"

	"github.com/brunobotter/site-sentinel/application/http"
	"github.com/brunobotter/site-sentinel/infra/logger"
	"github.com/brunobotter/site-sentinel/main/config"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type httpClient struct {
	cfg    *config.Config
	client *_http.Client
	logger logger.Logger
}

func NewHttpClient(cfg *config.Config, logger logger.Logger) http.Client {
	dialer := new(net.Dialer)

	transport := new(_http.Transport)
	transport.DialContext = dialer.DialContext
	transport.MaxIdleConnsPerHost = 5

	if cfg.Env == "local" {
		transport.Proxy = func(req *_http.Request) (*url.URL, error) {
			return _http.ProxyFromEnvironment(req)
		}
	}

	client := new(_http.Client)
	client.Transport = transport
	client = WrapClient(client)
	return &httpClient{cfg, client, logger}
}

func WrapClient(client *_http.Client) *_http.Client {
	return httptrace.WrapClient(client,
		httptrace.RTWithResourceNamer(func(req *_http.Request) string {
			return fmt.Sprintf("%s - %s", req.Method, req.Host)
		}),
	)
}

func (c *httpClient) NewRequestWithContext(ctx context.Context, method string, url string, body []byte) (request http.Request, err error) {
	request, err = newRequestWithContext(ctx, method, url, body)
	if err != nil {
		return request, err
	}
	return request, nil
}

func (c *httpClient) NewRequest(method string, url string, body []byte) (request http.Request, err error) {
	request, err = newRequest(method, url, body)
	if err != nil {
		return request, err
	}
	return request, nil
}

func (c *httpClient) Do(ctx context.Context, service string, req http.Request) (response http.Response, err error) {
	request, ok := req.(*request)
	if !ok {
		return response, errors.New("invalid request")
	}
	resp, err := c.client.Do(request.Req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, bodyReadErr := io.ReadAll(resp.Body)

	return newResponse(resp, body, bodyReadErr), nil
}
