package middlewares

/*import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobotter/site-sentinel/main/config"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestParseOrigins_Empty(t *testing.T) {
	assert.Nil(t, parseOrigins(""))
	assert.Nil(t, parseOrigins("   "))
}

func TestParseOrigins_Multiple(t *testing.T) {
	in := "https://a.example.com   https://b.example.com   https://c.example.com"
	out := parseOrigins(in)
	assert.Equal(t, []string{
		"https://a.example.com",
		"https://b.example.com",
		"https://c.example.com",
	}, out)
}

func TestNewCORSMiddleware_NoOrigins_BlockAll(t *testing.T) {
	e := echo.New()
	cfg := &config.Config{FrontURLs: ""}

	mw := NewCORSMiddleware(cfg)

	handlerCalled := false
	h := func(c echo.Context) error {
		handlerCalled = true
		return c.NoContent(http.StatusOK)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "https://any-origin.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.False(t, handlerCalled)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func TestBlockAllCORS_NoOrigin_PassThrough(t *testing.T) {
	e := echo.New()
	mw := blockAllCORS()

	handlerCalled := false
	h := func(c echo.Context) error {
		handlerCalled = true
		return c.String(http.StatusOK, "ok")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.True(t, handlerCalled)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestNewCORSMiddleware_WithOrigins_AllowsConfiguredOrigin(t *testing.T) {
	e := echo.New()
	cfg := &config.Config{FrontURLs: "https://allowed.example.com https://other.example.com"}

	mw := NewCORSMiddleware(cfg)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "https://allowed.example.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "ok", rec.Body.String())

	assert.Equal(t, "https://allowed.example.com", rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
	assert.Equal(t, "true", rec.Header().Get(echo.HeaderAccessControlAllowCredentials))
}

func TestNewCORSMiddleware_WithOrigins_DisallowedOrigin_NoCORSHeaders(t *testing.T) {
	e := echo.New()
	cfg := &config.Config{FrontURLs: "https://allowed.example.com"}

	mw := NewCORSMiddleware(cfg)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "https://not-allowed.example.com")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Empty(t, rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
	assert.Empty(t, rec.Header().Get(echo.HeaderAccessControlAllowCredentials))
}

func TestNewCORSMiddleware_PreflightAllowedOrigin(t *testing.T) {
	e := echo.New()
	cfg := &config.Config{FrontURLs: "https://allowed.example.com"}

	mw := NewCORSMiddleware(cfg)

	h := func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	}

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "https://allowed.example.com")
	req.Header.Set(echo.HeaderAccessControlRequestMethod, http.MethodGet)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "https://allowed.example.com", rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
	assert.Contains(t, rec.Header().Get(echo.HeaderAccessControlAllowMethods), http.MethodGet)
}
*/
