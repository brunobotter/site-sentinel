package logger

type jammesLoggerSetup struct {
	logger Logger
}

func initJammesLoggerSetup() *jammesLoggerSetup {
	return &jammesLoggerSetup{
		logger: NewJammesLogger("cartoes-web-api", "development", false),
	}
}
