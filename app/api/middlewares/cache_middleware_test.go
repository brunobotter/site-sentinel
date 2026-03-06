package middlewares

/*import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCacheMiddleware_GetMiddleware(t *testing.T) {
	middleware := &CacheMiddleware{
		middlewareFunc: "test-middleware",
	}

	result := middleware.GetMiddleware()

	assert.Equal(t, "test-middleware", result)
}

func TestNewCacheMiddleware_SetsCacheHeaders(t *testing.T) {
	e := echo.New()
	mw := NewCacheMiddleware()

	handlerCalled := false
	h := func(c echo.Context) error {
		handlerCalled = true
		return c.String(http.StatusOK, "success")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)

	assert.NoError(t, err)
	assert.True(t, handlerCalled)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "success", rec.Body.String())

	// Verify cache headers are set correctly
	assert.Equal(t, "no-cache, no-store, must-revalidate", rec.Header().Get("Cache-Control"))
	assert.Equal(t, "no-cache", rec.Header().Get("Pragma"))
	assert.Equal(t, "0", rec.Header().Get("Expires"))
}

func TestNewCacheMiddleware_CallsNextHandler(t *testing.T) {
	e := echo.New()
	mw := NewCacheMiddleware()

	nextHandlerCalled := false
	h := func(c echo.Context) error {
		nextHandlerCalled = true
		return nil
	}

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)

	assert.NoError(t, err)
	assert.True(t, nextHandlerCalled)
}

func TestNewCacheMiddleware_PropagatesHandlerError(t *testing.T) {
	e := echo.New()
	mw := NewCacheMiddleware()

	expectedError := echo.NewHTTPError(http.StatusInternalServerError, "test error")
	h := func(c echo.Context) error {
		return expectedError
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	// Headers should still be set even when handler returns error
	assert.Equal(t, "no-cache, no-store, must-revalidate", rec.Header().Get("Cache-Control"))
	assert.Equal(t, "no-cache", rec.Header().Get("Pragma"))
	assert.Equal(t, "0", rec.Header().Get("Expires"))
}

func TestNewCacheMiddleware_WithDifferentHTTPMethods(t *testing.T) {
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
	}

	for _, method := range methods {
		t.Run("method_"+method, func(t *testing.T) {
			e := echo.New()
			mw := NewCacheMiddleware()

			h := func(c echo.Context) error {
				return c.NoContent(http.StatusOK)
			}

			req := httptest.NewRequest(method, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := mw(h)(c)

			assert.NoError(t, err)
			assert.Equal(t, "no-cache, no-store, must-revalidate", rec.Header().Get("Cache-Control"))
			assert.Equal(t, "no-cache", rec.Header().Get("Pragma"))
			assert.Equal(t, "0", rec.Header().Get("Expires"))
		})
	}
}

func TestNewCacheMiddleware_HeadersSetBeforeHandler(t *testing.T) {
	e := echo.New()
	mw := NewCacheMiddleware()

	h := func(c echo.Context) error {
		// Verify headers are already set when handler is called
		assert.Equal(t, "no-cache, no-store, must-revalidate", c.Response().Header().Get("Cache-Control"))
		assert.Equal(t, "no-cache", c.Response().Header().Get("Pragma"))
		assert.Equal(t, "0", c.Response().Header().Get("Expires"))
		return c.NoContent(http.StatusOK)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := mw(h)(c)

	assert.NoError(t, err)
}
*/
