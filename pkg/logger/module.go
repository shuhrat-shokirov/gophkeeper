package logger

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gophkeeper/pkg/config"
)

var Module = fx.Provide(New)

type Logger interface {
	Info(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}

type Params struct {
	fx.In
	fx.Lifecycle

	Config config.Config
}

type logger struct {
	lg *zap.SugaredLogger
}

func New(p Params) (Logger, error) {
	level := getLevel(p.Config)

	developmentEncoderConfig := zap.NewDevelopmentEncoderConfig()

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(developmentEncoderConfig),
			zapcore.Lock(os.Stdout),
			level,
		))

	log := zap.New(core, zap.AddCaller())

	p.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				if err := log.Sync(); err != nil {
					return fmt.Errorf("failed to sync logger: %w", err)
				}
				return nil
			},
		},
	)

	return &logger{lg: log.Sugar()}, nil
}

func (l *logger) Info(msg string, fields ...interface{}) {
	l.lg.Infow(msg, fields...)
}

func (l *logger) Debug(msg string, fields ...interface{}) {
	l.lg.Debugw(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...interface{}) {
	l.lg.Warnw(msg, fields...)
}

func (l *logger) Error(msg string, fields ...interface{}) {
	l.lg.Errorw(msg, fields...)
}

func getLevel(config config.Config) zapcore.Level {
	switch config.GetString("logger.level") {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}
