package logger

import (
	"context"

	reqContext "github.com/Omelman/task-management/api/context"
	"go.uber.org/zap"
)

const requestIDKey = "request_id"

var logger *zap.Logger

func Get() *zap.Logger {
	return logger
}

func Load() (err error) {
	logger, err = zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:          "json",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: false,
	}.Build()
	return err
}

func WithCtxValue(ctx context.Context) *zap.Logger {
	return logger.With(zapFieldsFromContext(ctx)...)
}

func zapFieldsFromContext(ctx context.Context) []zap.Field {
	return []zap.Field{
		zap.String(requestIDKey, reqContext.GetRequestID(ctx)),
	}
}
