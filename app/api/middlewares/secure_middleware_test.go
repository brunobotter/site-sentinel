package middlewares

/*
import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSecureMiddleware_GetMiddleware(t *testing.T) {
	m := &SecureMiddleware{middlewareFunc: "test-func"}
	assert.Equal(t, "test-func", m.GetMiddleware())
}

func TestNewSecureMiddleware_ReturnsFunc(t *testing.T) {
	mw := NewSecureMiddleware()
	assert.NotNil(t, mw)
}

func TestNewSecureMiddleware_SetsSecurityHeaders_OnTLS(t *testing.T) {
	e := echo.New()
	mw := NewSecureMiddleware()

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.TLS = &tls.ConnectionState{}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "ok", rec.Body.String())

	assert.Equal(t, "1; mode=block", rec.Header().Get("X-XSS-Protection"))
	assert.Equal(t, "nosniff", rec.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "DENY", rec.Header().Get("X-Frame-Options"))
	sts := rec.Header().Get("Strict-Transport-Security")
	assert.Contains(t, sts, "max-age=31536000")
	assert.Contains(t, strings.ToLower(sts), "includesubdomains")
	assert.Equal(t, "strict-origin-when-cross-origin", rec.Header().Get("Referrer-Policy"))
}

func TestNewSecureMiddleware_NoHSTSOnNonTLS(t *testing.T) {
	e := echo.New()
	mw := NewSecureMiddleware()

	h := func(c echo.Context) error { return c.NoContent(http.StatusNoContent) }

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Empty(t, rec.Header().Get("Strict-Transport-Security"))
}
*/
