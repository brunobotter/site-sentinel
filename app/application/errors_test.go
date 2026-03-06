package application

/*
func TestNewIntegrationApplicationError(t *testing.T) {
	originalErr := errors.New("original error")
	err := NewIntegrationApplicationError(originalErr)

	assert.Error(t, err)
	assert.IsType(t, IntegrationApplicationError{}, errors.Unwrap(err))
	assert.Equal(t, "original error", errors.Unwrap(err).Error())
}

func TestNewTimeoutExceededApplicationError(t *testing.T) {
	originalErr := errors.New("timeout error")
	err := NewTimeoutExceededApplicationError(originalErr)

	assert.Error(t, err)
	assert.IsType(t, TimeoutExceededApplicationError{}, errors.Unwrap(err))
	assert.Equal(t, "timeout error", errors.Unwrap(err).Error())
}

func TestNewBadRequestApplicationError(t *testing.T) {
	originalErr := errors.New("bad request error")
	err := NewBadRequestApplicationError(originalErr)

	assert.Error(t, err)
	assert.IsType(t, BadRequestApplicationError{}, errors.Unwrap(err))
	assert.Equal(t, "bad request error", errors.Unwrap(err).Error())
}

func TestNewNotFoundApplicationError(t *testing.T) {
	originalErr := errors.New("not found error")
	err := NewNotFoundApplicationError(originalErr)

	assert.Error(t, err)
	assert.IsType(t, NotFoundApplicationError{}, errors.Unwrap(err))
	assert.Equal(t, "not found error", errors.Unwrap(err).Error())
}

func TestNewValidationApplicationError(t *testing.T) {
	originalErr := errors.New("validation error")
	err := NewValidationApplicationError(originalErr)

	assert.Error(t, err)
	assert.IsType(t, ValidationApplicationError{}, errors.Unwrap(err))
	assert.Equal(t, "validation error", errors.Unwrap(err).Error())
}

func TestNewForbiddenApplicationError(t *testing.T) {
	originalErr := errors.New("forbidden error")
	err := NewForbiddenApplicationError(originalErr)

	assert.Error(t, err)
	assert.IsType(t, ForbiddenApplicationError{}, errors.Unwrap(err))
	assert.Equal(t, "forbidden error", errors.Unwrap(err).Error())
}

func TestNewApplicationError(t *testing.T) {
	t.Run("Test nil error input", func(t *testing.T) {
		err := newApplicationError(NotFoundApplicationErrorType, nil)
		assert.Nil(t, err, "Expected nil error when input error is nil")
	})

	t.Run("Test ForbiddenApplicationErrorType", func(t *testing.T) {
		originalErr := errors.New("forbidden access")
		err := newApplicationError(ForbiddenApplicationErrorType, originalErr)

		assert.Error(t, err)
		assert.IsType(t, ForbiddenApplicationError{}, errors.Unwrap(err))
		assert.Equal(t, "forbidden access", errors.Unwrap(err).Error())
	})

	t.Run("Test default case", func(t *testing.T) {
		originalErr := errors.New("generic error")
		err := newApplicationError("UnknownErrorType", originalErr)

		assert.Error(t, err)
		assert.Equal(t, "generic error", errors.Unwrap(err).Error())
	})
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "additional context")

	if wrappedErr == nil {
		t.Fatal("expected non-nil error")
	}

	wrapped, ok := wrappedErr.(*WrappedError)
	if !ok {
		t.Fatalf("expected WrappedError, got %T", wrappedErr)
	}

	if wrapped.originalError != originalErr {
		t.Errorf("expected original error to be %v, got %v", originalErr, wrapped.originalError)
	}

	if len(wrapped.messages) != 1 || wrapped.messages[0] != "additional context" {
		t.Errorf("expected messages to contain 'additional context', got %v", wrapped.messages)
	}
}

func TestGetOriginalError(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "context")
	original := GetOriginalError(wrappedErr)

	if original != originalErr {
		t.Errorf("expected original error to be %v, got %v", originalErr, original)
	}
}

func TestWrappedError_Error(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := &WrappedError{
		originalError: originalErr,
		path:          "some/path",
		messages:      []string{"context1", "context2"},
	}

	expected := "some/path: context1; context2;  => original error"
	if wrappedErr.Error() != expected {
		t.Errorf("expected error string to be %q, got %q", expected, wrappedErr.Error())
	}
}
*/
