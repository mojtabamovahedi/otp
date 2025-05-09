package logger

import "go.uber.org/zap"

type Logger interface {
	Println(msg string)
	Fatal(err error)
}

type appLogger struct {
	logger *zap.Logger
}

func NewAppLogger(logger *zap.Logger) Logger {
	return &appLogger{logger: logger}
}

func (ap *appLogger) Println(msg string) {
	ap.logger.Info(msg)
}
func (ap *appLogger) Fatal(err error) {
	ap.logger.Fatal(err.Error())
}
