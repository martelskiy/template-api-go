package logger

import "go.uber.org/zap"

var sugaredLogger *zap.SugaredLogger

func init() {
	log, err := zap.NewProduction()
	if err != nil {
		panic("error initializing logger")
	}
	sugaredLogger = log.Sugar()
}

func Get() *zap.SugaredLogger {
	return sugaredLogger
}

func Dispose() {
	if sugaredLogger != nil {
		_ = sugaredLogger.Sync()
	}
}
