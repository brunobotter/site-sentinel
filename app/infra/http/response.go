package http

import (
	_http "net/http"

	httpContract "github.com/brunobotter/site-sentinel/application/http"
)

type response struct {
	resp        *_http.Response
	body        []byte
	bodyReadErr error
}

func newResponse(resp *_http.Response, body []byte, bodyReadErr error) httpContract.Response {
	return &response{resp, body, bodyReadErr}
}

func (s *response) Status() (status int) {
	return s.resp.StatusCode
}
func (s *response) Body() ([]byte, error) {
	if s.bodyReadErr != nil {
		return nil, s.bodyReadErr
	}
	return s.body, nil
}
func (s *response) Header(key string) map[string][]string {
	if key == "" {
		return s.resp.Header
	}

	if values, found := s.resp.Header[key]; found {
		return map[string][]string{key: values}
	}

	return map[string][]string{}
}
