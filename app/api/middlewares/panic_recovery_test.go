package middlewares

/*
import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func TestGetPanicRecoveryMiddlewareFunc(t *testing.T) {
	t.Run("should recover from panic and log error", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		mockLogger.On("Error", mock.Anything).Return().Twice()
		mockLogger.On("WithFields", mock.Anything).Return(mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			panic(errors.New("test panic"))
		})

		err := handler(echoCtx)

		assert.NoError(t, err)
		mockLogger.AssertExpectations(t)
		mockLogger.AssertNumberOfCalls(t, "Error", 2)
	})

	t.Run("should set datadog span tags on panic when span exists", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		mockLogger.On("Error", mock.Anything).Return().Times(2)
		mockLogger.On("WithFields", mock.Anything).Return(mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		span, ctx := tracer.StartSpanFromContext(ctx, "test.span")
		defer span.Finish()

		req = req.WithContext(ctx)
		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			panic("test panic with span")
		})

		_ = handler(echoCtx)

		mockLogger.AssertExpectations(t)
	})

	t.Run("should not panic when no panic occurs", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			return nil
		})

		err := handler(echoCtx)

		assert.NoError(t, err)
		mockLogger.AssertNotCalled(t, "Error", mock.Anything)
	})

	t.Run("should log request details on panic", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		path := "/test-path"
		method := http.MethodPost

		mockLogger.On("Error", mock.Anything).Return().Twice()
		mockLogger.On("WithFields", mock.Anything).Return(mockLogger)

		e := echo.New()
		req := httptest.NewRequest(method, path, nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			panic("test panic")
		})

		_ = handler(echoCtx)

		mockLogger.AssertExpectations(t)
		mockLogger.AssertNumberOfCalls(t, "Error", 2)
	})

	t.Run("should handle different panic types", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		mockLogger.On("Error", mock.Anything).Return().Times(2)
		mockLogger.On("WithFields", mock.Anything).Return(mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			panic(123)
		})

		err := handler(echoCtx)

		assert.NoError(t, err)
		mockLogger.AssertExpectations(t)
	})

	t.Run("should log stack trace on panic", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		var capturedErrors []error
		mockLogger.On("Error", mock.Anything).Run(func(args mock.Arguments) {
			capturedErrors = append(capturedErrors, args.Get(0).(error))
		}).Return().Twice()
		mockLogger.On("WithFields", mock.Anything).Return(mockLogger)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			panic("test panic")
		})

		_ = handler(echoCtx)

		assert.Len(t, capturedErrors, 2)
		assert.Contains(t, capturedErrors[0].Error(), "PANIC RECOVERED IN REQUEST")
		assert.Contains(t, capturedErrors[1].Error(), "stack trace")
		mockLogger.AssertExpectations(t)
	})

	t.Run("should call next handler when no panic", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		nextCalled := false

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)

		ctx := logger.SetContextLogger(req.Context(), mockLogger)
		req = req.WithContext(ctx)

		rec := httptest.NewRecorder()
		echoCtx := e.NewContext(req, rec)

		middleware := getPanicRecoveryMiddlewareFunc(mockLogger)
		handler := middleware(func(c echo.Context) error {
			nextCalled = true
			return nil
		})

		err := handler(echoCtx)

		assert.NoError(t, err)
		assert.True(t, nextCalled)
		mockLogger.AssertNotCalled(t, "Error", mock.Anything)
	})
}

func TestNewPanicMiddleware(t *testing.T) {
	t.Run("should create panic middleware with logger", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)

		middleware := NewPanicMiddleware(mockLogger)

		assert.NotNil(t, middleware)
		assert.NotNil(t, middleware.GetMiddleware())
		assert.IsType(t, &PanicMiddleware{}, middleware)
	})
}

func TestPanicMiddleware_GetMiddleware(t *testing.T) {
	t.Run("should return middleware function", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		panicMiddleware := NewPanicMiddleware(mockLogger)

		middlewareFunc := panicMiddleware.GetMiddleware()

		assert.NotNil(t, middlewareFunc)
		assert.IsType(t, echo.MiddlewareFunc(nil), middlewareFunc)
	})

	t.Run("should return same middleware function on multiple calls", func(t *testing.T) {
		mockLogger := new(mocks.MockLogger)
		panicMiddleware := NewPanicMiddleware(mockLogger)

		middlewareFunc1 := panicMiddleware.GetMiddleware()
		middlewareFunc2 := panicMiddleware.GetMiddleware()

		assert.Equal(t, fmt.Sprintf("%p", middlewareFunc1), fmt.Sprintf("%p", middlewareFunc2))
	})
}
*/
