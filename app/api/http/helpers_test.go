package http

/*func TestHandleError_NotFound(t *testing.T) {
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	err := appErr.NewNotFoundApplicationError(errors.New("not found"))
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, "not found", resp.ErrorMessage)
}

func TestHandleError_TimeoutExceeded(t *testing.T) {
	err := appErr.NewTimeoutExceededApplicationError(errors.New("timeout"))
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 504, resp.StatusCode)
	assert.Equal(t, "timeout", resp.ErrorMessage)
}

func TestHandleError_Integration(t *testing.T) {
	err := appErr.NewIntegrationApplicationError(errors.New("integration fail"))
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 502, resp.StatusCode)
	assert.Equal(t, "integration fail", resp.ErrorMessage)
}

func TestHandleError_BadRequest(t *testing.T) {
	err := appErr.NewBadRequestApplicationError(errors.New("bad request"))
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "bad request", resp.ErrorMessage)
}

func TestHanderError_BadGateway(t *testing.T) {
	err := appErr.NewIntegrationApplicationError(errors.New("bad gateway"))
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 502, resp.StatusCode)
	assert.Equal(t, "bad gateway", resp.ErrorMessage)
}

func TestHandleError_Validation(t *testing.T) {
	err := appErr.NewValidationApplicationError(errors.New("validation fail"))
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 422, resp.StatusCode)
	assert.Equal(t, "validation fail", resp.ErrorMessage)
}

func TestHandleError_Forbidden(t *testing.T) {
	err := appErr.NewForbiddenApplicationError(errors.New("forbidden access"))
	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 403, resp.StatusCode)
	assert.Equal(t, "forbidden access", resp.ErrorMessage)
}

func TestHandleError_InternalServerError(t *testing.T) {
	err := errors.New("generic error")

	builder := mocks.NewSetup().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	resp := HandleError(context.Background(), err, builder.Logger)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, "generic error", resp.ErrorMessage)
}

func TestSpanFromContextManualMock(t *testing.T) {
	builder := mocks.NewSetup().WithDatadog().WithLogger()
	builder.Logger.On("Errorf", "Error handled: %s", mock.Anything).Return()
	dataDogMock := builder.Datadog
	type contextKey string
	const spanKey contextKey = "span"
	ctx := context.WithValue(context.Background(), spanKey, dataDogMock)
	errorToHandle := errors.New("test error")

	HandleError(ctx, errorToHandle, builder.Logger)

	spanFromContext, ok := ctx.Value(spanKey).(*mocks.MockSpan)
	if ok {
		spanFromContext.SetTag("error", true)
		spanFromContext.SetTag("error.message", errorToHandle.Error())
	}

	if dataDogMock.Tag("error") != true {
		t.Errorf("expected error tag to be true, got %v", dataDogMock.Tag("error"))
	}
	if dataDogMock.Tag("error.message") != "test error" {
		t.Errorf("expected error.message tag to be 'test error', got %v", dataDogMock.Tag("error.message"))
	}
}
*/
