package logger

import (
	"context"
	"fmt"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	JammesLoggerKey = "JammesLogger"
)

type jammes struct {
	appName      string
	level        string
	logger       *zap.Logger
	commonFields []any
}

func NewJammesLogger(appName string, environment string, debug bool) Logger {
	var config zap.Config
	var zapLogger *zap.Logger

	if environment == "production" {
		config = zap.NewProductionConfig()
		config.Encoding = "json"
		config.EncoderConfig = buildEncondingConfig()
		config.DisableStacktrace = true
		config.DisableCaller = true
	} else {
		config = zap.NewDevelopmentConfig()
		config.Encoding = "json"
		config.EncoderConfig = buildEncondingConfig()

	}

	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	if debug {
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	zapLogger, _ = config.Build()

	j := &jammes{
		appName:      appName,
		level:        config.Level.String(),
		logger:       zapLogger,
		commonFields: []any{},
	}

	j.SetCommonFields(map[string]any{
		"application_name": appName,
	})

	return j
}

func (j *jammes) SetCommonFields(commonFields map[string]any) {
	for key, value := range commonFields {
		j.commonFields = append(j.commonFields, zap.Any(key, value))
	}
}

func (j *jammes) Log(msg string) {
	j.Info(msg)
}

func (j *jammes) Print(message string) string {
	return message
}

func buildEncondingConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		TimeKey:        "timestamp",
		FunctionKey:    zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func (j *jammes) hideSensitiveData(data interface{}) (interface{}, error) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the provided value is not a struct: '%v'", val.Kind())
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := val.Type().Field(i)

		tag := typeField.Tag.Get("jammes")

		if tag == "hide" {
			if field.CanSet() {
				switch field.Kind() {
				case reflect.String:
					field.SetString("******")
				}
			}
		}

		if field.Kind() == reflect.Struct {
			j.hideSensitiveData(field.Addr().Interface())
		} else if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			j.hideSensitiveData(field.Interface())
		}
	}

	return data, nil
}

func (j *jammes) Infof(format string, args ...interface{}) {
	defer j.logger.Sync()

	j.logger.Sugar().With(j.commonFields...).Infof(format, args...)
}

func (j *jammes) Info(args ...interface{}) {
	defer j.logger.Sync()

	j.logger.Sugar().With(j.commonFields...).Info(args...)
}

func (j *jammes) WithFields(fields map[string]any) Logger {
	newLogger := &jammes{
		appName:      j.appName,
		level:        j.level,
		logger:       j.logger,
		commonFields: append([]any{}, j.commonFields...),
	}

	for key, value := range fields {
		newLogger.commonFields = append(newLogger.commonFields, zap.Any(key, value))
	}

	return newLogger
}

func (j *jammes) WithContext(ctx context.Context) Logger {
	fields := map[string]any{}
	return j.WithFields(fields)
}

func (j *jammes) Errorf(format string, args ...interface{}) {
	defer j.logger.Sync()

	j.logger.Sugar().With(j.commonFields...).Errorf(format, args...)
}

func (j *jammes) Error(err error) {
	defer j.logger.Sync()

	j.logger.Sugar().With(j.commonFields...).Error(err)
}
func (j *jammes) Debugf(format string, args ...interface{}) {
	defer j.logger.Sync()

	j.logger.Sugar().With(j.commonFields...).Debugf(format, args...)
}

func (j *jammes) Debug(args ...interface{}) {
	defer j.logger.Sync()

	j.logger.Sugar().With(j.commonFields...).Debug(args...)
}
