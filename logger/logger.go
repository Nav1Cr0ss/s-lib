package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

// NewLogger return logger instance
func NewLogger(debug bool) *Logger {
	loggers := map[bool]func(options ...zap.Option) (*zap.Logger, error){
		true:  zap.NewDevelopment,
		false: zap.NewProduction,
	}

	logger, _ := loggers[debug]()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	return &Logger{sugar}
}
