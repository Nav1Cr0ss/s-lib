package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
	CL *zap.Logger
}

// NewLogger return logger instance
func NewLogger(debug bool) *Logger {

	configs := map[bool]func() zap.Config{
		true:  zap.NewDevelopmentConfig,
		false: zap.NewProductionConfig,
	}

	config := configs[debug]()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stderr"}
	legacyLogger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	defer legacyLogger.Sync() // flushes buffer, if any

	sugarLogger := legacyLogger.Sugar()

	return &Logger{
		SugaredLogger: sugarLogger,
		CL:            legacyLogger,
	}
}
