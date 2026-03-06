package middlewares

/*
func TestLoggerMiddleware(t *testing.T) {
	e := echo.New()
	builder := mocks.NewSetup().WithLogger().WithConfig()
	builder.Logger.On("Infof", mock.Anything, mock.Anything)
	builder.Logger.On("WithFields", mock.Anything).Return(builder.Logger)
	builder.Logger.On("WithContext", mock.Anything).Return(builder.Logger)
	builder.Logger.On("WithFields", mock.Anything).Return(builder.Logger)

	middleware := getLoggerMiddlewareFunc(builder.Logger, builder.Config)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}

	t.Run("Middleware Execution", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := middleware(handler)(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	})

	t.Run("Middleware Execution in user agent empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := middleware(handler)(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	})

	t.Run("IsFromMobile - Mobile User Agent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		isMobile := IsFromMobile(c, req.Header.Get("user-agent"))
		assert.True(t, isMobile, "Expected IsFromMobile to return true for mobile User-Agent")
	})

	t.Run("IsFromMobile - Non-Mobile User Agent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		isMobile := IsFromMobile(c, req.Header.Get("user-agent"))
		assert.False(t, isMobile, "Expected IsFromMobile to return false for non-mobile User-Agent")
	})

	t.Run("IsFromMobile - Empty User Agent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		isMobile := IsFromMobile(c, req.Header.Get("user-agent"))
		assert.False(t, isMobile, "Expected IsFromMobile to return false for empty User-Agent")
	})
}

func TestGetMiddleware(t *testing.T) {
	builder := mocks.NewSetup().WithLogger().WithConfig()
	loggerMiddleware := &LoggerMiddleware{
		middlewareFunc: getLoggerMiddlewareFunc(builder.Logger, builder.Config),
	}

	result := loggerMiddleware.GetMiddleware()
	assert.NotNil(t, result, "GetMiddleware should return a non-nil middleware function")
}

func TestNewLoggerMiddleware(t *testing.T) {
	builder := mocks.NewSetup().WithLogger().WithConfig()

	result := NewLoggerMiddleware(builder.Logger, builder.Config)

	loggerMiddleware, ok := result.(*LoggerMiddleware)
	assert.True(t, ok, "NewLoggerMiddleware should return an instance of LoggerMiddleware")
	assert.NotNil(t, loggerMiddleware.middlewareFunc, "LoggerMiddleware should have a non-nil middlewareFunc")
}
*/
