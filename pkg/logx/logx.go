package logx

import (
	"context"
	"log/slog"
)

type logger struct {
	l *slog.Logger
}

var (
	Debug func(context.Context, string, ...any)
	Info func(context.Context, string, ...any)
	Warn func(context.Context, string, ...any)
	Error func(context.Context, string, ...any)
)

func Register(l logger) {
	// TODO
}