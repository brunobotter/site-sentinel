package shared

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type ContextKey string

const (
	ContextApplicationName        ContextKey = "application_name"
	ContextApplicationDebug       ContextKey = "application_debug"
	ContextApplicationEnvironment ContextKey = "application_environment"
	ContextJourneyID              ContextKey = "journey_id"
	ContextClientID               ContextKey = "client_id"
	UserAgentKey                  ContextKey = "user_agent"
	IsFromMobileKey               ContextKey = "is_from_mobile"
	ContextCorrelationID          ContextKey = "correlation_id"
	SessionID                     ContextKey = "session_id"
	CardID                        ContextKey = "card_id"
	JourneyName                   ContextKey = "journey_name"
	Ip                            ContextKey = "ip"
	Referer                       ContextKey = "referer"
	RecaptchaToken                ContextKey = "x-itau-recaptcha-token"
)

type ApplicationConfig struct {
	ApplicationName string
	Debug           bool
	Environment     string
}

func SetContextApplicationName(ctx context.Context, appName string) context.Context {
	return context.WithValue(ctx, ContextApplicationName, appName)
}

func SetContextApplicationDebug(ctx context.Context, appDebug bool) context.Context {
	return context.WithValue(ctx, ContextApplicationDebug, appDebug)
}

func SetContextApplicationEnvironment(ctx context.Context, environment string) context.Context {
	return context.WithValue(ctx, ContextApplicationEnvironment, environment)
}

func SetContextClientID(ctx context.Context, clientID string) context.Context {
	return context.WithValue(ctx, ContextClientID, clientID)
}

func SetCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, ContextCorrelationID, correlationID)
}

func SetSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionID, sessionID)
}

func SetJourneyID(ctx context.Context, journeyID string) context.Context {
	return context.WithValue(ctx, ContextJourneyID, journeyID)
}

func SetUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, UserAgentKey, userAgent)
}

func SetIsFromMobile(ctx context.Context, isFromMobile bool) context.Context {
	return context.WithValue(ctx, IsFromMobileKey, isFromMobile)
}

func SetCardID(ctx context.Context, cardID string) context.Context {
	return context.WithValue(ctx, CardID, cardID)
}

func SetJourneyName(ctx context.Context, journeyName string) context.Context {
	return context.WithValue(ctx, JourneyName, journeyName)
}

func SetIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, Ip, ip)
}

func SetRecaptchaToken(ctx context.Context, recaptchatoken string) context.Context {
	return context.WithValue(ctx, RecaptchaToken, recaptchatoken)
}

func SetReferer(ctx context.Context, referer string) context.Context {
	return context.WithValue(ctx, Referer, referer)
}

func FromContextApplicationConfig(ctx context.Context) ApplicationConfig {
	appName := ctx.Value(ContextApplicationName)
	appDebug := ctx.Value(ContextApplicationDebug)
	appEnvironment := ctx.Value(ContextApplicationEnvironment)

	debug, _ := strconv.ParseBool(fmt.Sprint(appDebug))

	return ApplicationConfig{
		ApplicationName: fmt.Sprint(appName),
		Debug:           debug,
		Environment:     fmt.Sprint(appEnvironment),
	}
}

func GetJourneyIDFromContext(ctx context.Context) string {
	if journeyID, ok := ctx.Value(ContextJourneyID).(string); ok {
		return journeyID
	}
	return ""
}

func GetSessionIDFromContext(ctx context.Context) string {
	if sessionID, ok := ctx.Value(SessionID).(string); ok {
		return sessionID
	}
	return ""
}

func GetCorrelationIDFromContext(ctx context.Context) string {
	if correlationId, ok := ctx.Value(ContextCorrelationID).(string); ok {
		return correlationId
	}
	return uuid.NewString()
}

func GetCardIDFromContext(ctx context.Context) string {
	if cardID, ok := ctx.Value(CardID).(string); ok {
		return cardID
	}
	return ""
}

func GetJourneyNameFromContext(ctx context.Context) string {
	if journeyName, ok := ctx.Value(JourneyName).(string); ok {
		return journeyName
	}
	return ""
}

func GetIPFromContext(ctx context.Context) string {
	if ip, ok := ctx.Value(Ip).(string); ok {
		return ip
	}
	return ""
}

func GetUserAgentFromContext(ctx context.Context) string {
	if userAgent, ok := ctx.Value(UserAgentKey).(string); ok {
		return userAgent
	}
	return ""
}

func GetRecaptchaTokenFromContext(ctx context.Context) string {
	if recaptchaToken, ok := ctx.Value(RecaptchaToken).(string); ok {
		return recaptchaToken
	}
	return ""
}

func IsAndroid(userAgent string) bool {
	return strings.Contains(strings.ToLower(userAgent), "android")
}

func GetRefererFromContext(ctx context.Context) string {
	if referer, ok := ctx.Value(Referer).(string); ok {
		return referer
	}
	return ""
}
