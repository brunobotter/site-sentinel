package shared

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetContextApplicationName(t *testing.T) {
	ctx := context.Background()
	ctx = SetContextApplicationName(ctx, "TestApp")

	appName := ctx.Value(ContextApplicationName)
	assert.Equal(t, "TestApp", appName)
}

func TestSetContextApplicationDebug(t *testing.T) {
	ctx := context.Background()
	ctx = SetContextApplicationDebug(ctx, true)

	appDebug := ctx.Value(ContextApplicationDebug)
	assert.Equal(t, true, appDebug)
}

func TestSetContextApplicationEnvironment(t *testing.T) {
	ctx := context.Background()
	ctx = SetContextApplicationEnvironment(ctx, "production")

	appEnvironment := ctx.Value(ContextApplicationEnvironment)
	assert.Equal(t, "production", appEnvironment)
}

func TestSetContextClientID(t *testing.T) {
	ctx := context.Background()
	ctx = SetContextClientID(ctx, "client-456")

	clientID := ctx.Value(ContextClientID)
	assert.Equal(t, "client-456", clientID)
}

func TestFromContextApplicationConfig(t *testing.T) {
	ctx := context.Background()
	ctx = SetContextApplicationName(ctx, "TestApp")
	ctx = SetContextApplicationDebug(ctx, true)
	ctx = SetContextApplicationEnvironment(ctx, "production")
	config := FromContextApplicationConfig(ctx)
	assert.Equal(t, "TestApp", config.ApplicationName)
	assert.Equal(t, true, config.Debug)
	assert.Equal(t, "production", config.Environment)
}

func TestSetContextCorrelationID(t *testing.T) {
	ctx := context.Background()
	ctx = SetCorrelationID(ctx, "123")

	correlationID := ctx.Value(ContextCorrelationID)
	assert.Equal(t, "123", correlationID)
}

func TestSetContextUserAgent(t *testing.T) {
	ctx := context.Background()
	ctx = SetUserAgent(ctx, "mobile 123")

	userAgent := ctx.Value(UserAgentKey)
	assert.Equal(t, "mobile 123", userAgent)
}

func TestSetContextFromMobile(t *testing.T) {
	ctx := context.Background()
	ctx = SetIsFromMobile(ctx, true)

	fromMobile := ctx.Value(IsFromMobileKey)
	assert.Equal(t, true, fromMobile)
}

func TestGetJourneyIDFromContext(t *testing.T) {
	t.Run("should return journeyID when present in context", func(t *testing.T) {
		expectedJourneyID := "12345"
		ctx := context.WithValue(context.Background(), ContextJourneyID, expectedJourneyID)
		result := GetJourneyIDFromContext(ctx)
		if result != expectedJourneyID {
			t.Errorf("expected %s, got %s", expectedJourneyID, result)
		}
	})

	t.Run("should return empty string when journeyID is not present in context", func(t *testing.T) {
		ctx := context.Background()
		result := GetJourneyIDFromContext(ctx)
		if result != "" {
			t.Errorf("expected empty string, got %s", result)
		}
	})

	t.Run("should return empty string when journeyID is not a string", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), ContextJourneyID, 12345)
		result := GetJourneyIDFromContext(ctx)
		if result != "" {
			t.Errorf("expected empty string, got %s", result)
		}
	})
}
