package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestSetCommonFields(t *testing.T) {
	jammesLogger := initJammesLoggerSetup()

	commonFields := map[string]any{
		"lead_uuid": "abcd1234",
	}

	jammesLogger.logger.SetCommonFields(commonFields)

	expectedCommonFields := []any{
		zapcore.Field{
			Key:       "application_name",
			String:    "cartoes-web-api",
			Type:      15,
			Integer:   0,
			Interface: nil,
		},
		zapcore.Field{
			Key:       "lead_uuid",
			Type:      15,
			Integer:   0,
			String:    "abcd1234",
			Interface: nil,
		},
	}

	assert.Equal(t, expectedCommonFields, jammesLogger.logger.(*jammes).commonFields)
}

func TestPrint(t *testing.T) {
	jammeLogger := initJammesLoggerSetup()

	output := jammeLogger.logger.Print("test message")

	assert.Equal(t, "test message", output)
}

func TestHideSensitiveDataWithInvalidType(t *testing.T) {
	jammesLogger := initJammesLoggerSetup()

	dataToHide := "invalid"

	hiddenData, err := jammesLogger.logger.(*jammes).hideSensitiveData(&dataToHide)

	assert.NotNil(t, err)
	assert.Error(t, err, "the provided value is not a struct: 'string'")
	assert.Nil(t, hiddenData)
}
