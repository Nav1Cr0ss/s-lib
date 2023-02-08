package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

// NewLogger return logger instance
func NewLogger(debug bool) *Logger {
	configs := map[bool]func() zap.Config{
		true:  zap.NewDevelopmentConfig,
		false: zap.NewProductionConfig,
	}

	config := configs[debug]()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger, _ := config.Build()
	defer zapLogger.Sync() // flushes buffer, if any

	sugar := zapLogger.Sugar()

	return &Logger{sugar}
}
