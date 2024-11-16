package logging

import (
	"fmt"
	"log"

	"github.com/JeongJaeSoon/go-auth/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(cfg config.LoggingConfig) (*zap.Logger, error) {
	logLevel := zap.InfoLevel

	if err := logLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		logLevel = zap.InfoLevel
		log.Printf("Invalid log level in config, using default level: %v", logLevel)
	}

	zapConfig := zap.Config{
		Encoding:         cfg.Encoding,
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %v", err)
	}

	return logger, nil
}
