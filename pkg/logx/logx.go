package logx

import (
	"context"
	"log/slog"

	defaultlogger "github.com/TamasGorgics/gomag/pkg/logx/internal/default"
)

type Logger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Warn(context.Context, string, ...any)
	Error(context.Context, error, string, ...any)
	Fatal(context.Context, error, string, ...any)
	With(ctx context.Context, key string, val any) context.Context
}

// Global variables for convenience
var (
	Debug func(context.Context, string, ...any)
	Info  func(context.Context, string, ...any)
	Warn  func(context.Context, string, ...any)
	Error func(context.Context, error, string, ...any)
	Fatal func(context.Context, error, string, ...any)
	With  func(ctx context.Context, key string, val any) context.Context
)

func register(l Logger) {
	if l == nil {
		return
	}

	Debug = l.Debug
	Info = l.Info
	Warn = l.Warn
	Error = l.Error
	Fatal = l.Fatal
	With = l.With
}

func InitDefaultLogger() Logger {
	l := defaultlogger.NewDefaultLogger(slog.LevelInfo)
	register(l)
	return l
}
