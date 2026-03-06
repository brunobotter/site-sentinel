package logger

import (
	"context"

	"github.com/brunobotter/site-sentinel/util/shared"
)

const (
	ContextLoggerKey shared.ContextKey = "logger"
)

func SetContextLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ContextLoggerKey, logger)
}

func getLoggerFromContext(ctx context.Context) Logger {
	applicationConfig := shared.FromContextApplicationConfig(ctx)

	applicationName := applicationConfig.ApplicationName
	applicationDebug := applicationConfig.Debug
	applicationEnvironment := applicationConfig.Environment

	jammesLogger, _ := ctx.Value(ContextLoggerKey).(Logger)

	if jammesLogger == nil {
		jammesLogger = NewJammesLogger(applicationName, applicationEnvironment, applicationDebug)
	}

	sessionID := shared.GetSessionIDFromContext(ctx)
	journeyID := shared.GetJourneyIDFromContext(ctx)
	correlationID := shared.GetCorrelationIDFromContext(ctx)
	cardID := shared.GetCardIDFromContext(ctx)
	journeyName := shared.GetJourneyNameFromContext(ctx)

	tags := make(map[string]any)
	if sessionID != "" {
		tags["session_id"] = sessionID
	}
	if journeyID != "" {
		tags["journey_id"] = journeyID
	}

	if cardID != "" {
		tags["card_id"] = cardID
	}

	if journeyName != "" {
		tags["journey_name"] = journeyName
	}

	jammesLogger = jammesLogger.WithFields(tags)

	// envia as tags para o datadog
	ip := shared.GetIPFromContext(ctx)
	userAgent := shared.GetUserAgentFromContext(ctx)
	if ip != "" {
		tags["ip"] = ip
	}
	if userAgent != "" {
		tags["user_agent"] = userAgent
	}
	if correlationID != "" {
		tags["correlation_id"] = correlationID
	}

	return jammesLogger
}

func LoggerFromContext(ctx context.Context) Logger {
	return getLoggerFromContext(ctx)
}
