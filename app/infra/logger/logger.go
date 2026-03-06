package logger

import "context"

type Logger interface {
	SetCommonFields(commonFields map[string]any)
	Print(message string) string
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	WithFields(fields map[string]any) Logger
	Errorf(format string, args ...interface{})
	Error(err error)
	WithContext(ctx context.Context) Logger
	Log(msg string)
}
