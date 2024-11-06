package logging

import (
	"log"
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

		var err error
		Logger, err = zapConfig.Build()
		if err != nil {
			panic("failed to build logger: " + err.Error())
		}
	})
}
