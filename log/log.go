package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger = nil

func New() {
	// Allow initialization only once
	if log == nil {
		logConfig := zap.Config{
			OutputPaths: []string{"stdout"},
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				LevelKey:         "level",
				TimeKey:          "time",
				MessageKey:       "msg",
				EncodeTime:       zapcore.ISO8601TimeEncoder,
				EncodeLevel:      zapcore.LowercaseLevelEncoder,
				EncodeCaller:     zapcore.ShortCallerEncoder,
				ConsoleSeparator: " ",
			},
		}

		var err error
		if log, err = logConfig.Build(); err != nil {
			panic(err)
		}
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("ERROR", err))
	}

	log.Error(msg, tags...)
	log.Sync()
}

func Fatal(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("FATAL", err))
	}

	log.Fatal(msg, tags...)
	log.Sync()
}
