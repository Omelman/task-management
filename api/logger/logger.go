package logger

import (
	"go.uber.org/zap"
)

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
