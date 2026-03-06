package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	body := map[string]string{"foo": "bar"}
	resp := Ok(body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, HttpResponseData{body}, resp.Body)
	assert.Empty(t, resp.ErrorMessage)
}

func TestOkNoContent(t *testing.T) {
	resp := OkNoContent()
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Nil(t, resp.Body)
	assert.Empty(t, resp.ErrorMessage)
}

func TestCreated(t *testing.T) {
	body := map[string]string{"foo": "bar"}
	resp := Created(body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, HttpResponseData{body}, resp.Body)
	assert.Empty(t, resp.ErrorMessage)
}

func TestCreatedNoContent(t *testing.T) {
	resp := CreatedNoContent()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Nil(t, resp.Body)
	assert.Empty(t, resp.ErrorMessage)
}

func TestUnauthorized(t *testing.T) {
	msg := "unauthorized access"
	resp := Unauthorized(msg)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestForbidden(t *testing.T) {
	msg := "forbidden access"
	resp := Forbidden(msg)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestUnprocessableEntity(t *testing.T) {
	msg := "unprocessable entity"
	resp := UnprocessableEntity(msg)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestNotFound(t *testing.T) {
	msg := "not found"
	resp := NotFound(msg)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestServiceUnavailable(t *testing.T) {
	msg := "service unavailable"
	resp := ServiceUnavailable(msg)
	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestInternalServerError(t *testing.T) {
	msg := "internal error"
	resp := InternalServerError(msg)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestTimeoutExceeded(t *testing.T) {
	msg := "timeout"
	resp := TimeoutExceeded(msg)
	assert.Equal(t, http.StatusGatewayTimeout, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestBadRequest(t *testing.T) {
	msg := "bad request"
	resp := BadRequest(msg)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestBadGateway(t *testing.T) {
	msg := "bad gateway"
	resp := BadGateway(msg)
	assert.Equal(t, http.StatusBadGateway, resp.StatusCode)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, resp.Body)
	assert.Equal(t, msg, resp.ErrorMessage)
}

func TestBuildErrorBody(t *testing.T) {
	msg := "some error"
	body := buildErrorBody(msg)
	assert.Equal(t, map[string]any{"error": true, "message": msg}, body)
}
