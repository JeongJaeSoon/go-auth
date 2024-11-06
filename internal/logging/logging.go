package logging

import (
	"sync"

	"github.com/JeongJaeSoon/go-auth/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	once   sync.Once
)

func InitLogger(cfg config.LoggingConfig) {
	once.Do(func() {
		logLevel := zap.InfoLevel

		if err := logLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
			logLevel = zap.InfoLevel
		}

		zapConfig := zap.Config{
			Encoding:         cfg.Format,
			Level:            zap.NewAtomicLevelAt(logLevel),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "timestamp",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
			},
		}

		var err error
		Logger, err = zapConfig.Build()
		if err != nil {
			panic("failed to build logger: " + err.Error())
		}
	})
}
