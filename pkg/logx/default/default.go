package defaultx

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/TamasGorgics/gomag/pkg/logx"
)

// defaultLogger is a logger that uses slog under the hood
type defaultLogger struct {
	level logx.Level
}

func NewDefaultLogger(level logx.Level) logx.Logger {
	return &defaultLogger{
		level: level,
	}
}

func (l *defaultLogger) Debug(ctx context.Context, format string, args ...any) {
	slog.DebugContext(ctx, format, args...)
}

func (l *defaultLogger) Info(ctx context.Context, format string, args ...any) {
	slog.InfoContext(ctx, format, args...)
}

func (l *defaultLogger) Warn(ctx context.Context, format string, args ...any) {
	slog.WarnContext(ctx, format, args...)
}

func (l *defaultLogger) Error(ctx context.Context, err error, format string, args ...any) {
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("%s (%s)", err.Error(), fmt.Sprintf(format, args...)), nil)
	} else {
		slog.ErrorContext(ctx, format, args...)
	}
}

func (l *defaultLogger) Fatal(ctx context.Context, err error, format string, args ...any) {
	l.Error(ctx, err, format, args...)
	panic(format)
}

func (l *defaultLogger) With(ctx context.Context, key string, val any) context.Context {
	return context.WithValue(ctx, key, val)
}
