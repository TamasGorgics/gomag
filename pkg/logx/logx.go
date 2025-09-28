package logx

import (
	"context"
)

type Logger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Warn(context.Context, string, ...any)
	Error(context.Context, error, string, ...any)
	Fatal(context.Context, error, string, ...any)
	With(ctx context.Context, key string, val any) context.Context
}

var (
	Debug func(context.Context, string, ...any)
	Info  func(context.Context, string, ...any)
	Warn  func(context.Context, string, ...any)
	Error func(context.Context, error, string, ...any)
	Fatal func(context.Context, error, string, ...any)
	With  func(ctx context.Context, key string, val any) context.Context
)

func Register(l Logger) {
	Debug = l.Debug
	Info = l.Info
	Warn = l.Warn
	Error = l.Error
	Fatal = l.Fatal
	With = l.With
}
