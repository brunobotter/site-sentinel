package http

import (
	"errors"
	_http "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseStatus(t *testing.T) {
	resp := newResponse(&_http.Response{StatusCode: _http.StatusCreated}, []byte("ok"), nil)

	assert.Equal(t, _http.StatusCreated, resp.Status())
}

func TestResponseBody(t *testing.T) {
	resp := newResponse(&_http.Response{}, []byte("payload"), nil)

	body, err := resp.Body()

	assert.NoError(t, err)
	assert.Equal(t, []byte("payload"), body)
}

func TestResponseBodyError(t *testing.T) {
	readErr := errors.New("failed to read body")
	resp := newResponse(&_http.Response{}, nil, readErr)

	body, err := resp.Body()

	assert.Nil(t, body)
	assert.Equal(t, readErr, err)
}

func TestResponseHeaderWithEmptyKeyReturnsAllHeaders(t *testing.T) {
	headers := _http.Header{
		"Content-Type": {"application/json"},
		"X-Test":       {"value"},
	}
	resp := newResponse(&_http.Response{Header: headers}, nil, nil)

	assert.Equal(t, map[string][]string(headers), resp.Header(""))
}

func TestResponseHeaderWithSpecificKeyReturnsRequestedHeader(t *testing.T) {
	headers := _http.Header{
		"Content-Type": {"application/json"},
		"X-Test":       {"value"},
	}
	resp := newResponse(&_http.Response{Header: headers}, nil, nil)

	assert.Equal(t, map[string][]string{"X-Test": {"value"}}, resp.Header("X-Test"))
}

func TestResponseHeaderWhenKeyDoesNotExistReturnsEmptyMap(t *testing.T) {
	resp := newResponse(&_http.Response{Header: _http.Header{}}, nil, nil)

	assert.Equal(t, map[string][]string{}, resp.Header("X-Missing"))
}
