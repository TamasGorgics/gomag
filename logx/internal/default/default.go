package defaultlogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// defaultLogger is a logger that uses slog under the hood
type DefaultLogger struct {
	level  slog.Level
	logger *slog.Logger
}

func NewDefaultLogger(level slog.Level) *DefaultLogger {
	return &DefaultLogger{
		level:  level,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})),
	}
}

func (l *DefaultLogger) Debug(ctx context.Context, format string, args ...any) {
	l.logger.DebugContext(ctx, format, args...)
}

func (l *DefaultLogger) Info(ctx context.Context, format string, args ...any) {
	l.logger.InfoContext(ctx, format, args...)
}

func (l *DefaultLogger) Warn(ctx context.Context, format string, args ...any) {
	l.logger.WarnContext(ctx, format, args...)
}

func (l *DefaultLogger) Error(ctx context.Context, err error, format string, args ...any) {
	if err != nil {
		l.logger.ErrorContext(ctx, fmt.Sprintf("%s (%s)", err.Error(), fmt.Sprintf(format, args...)))
	} else {
		l.logger.ErrorContext(ctx, format, args...)
	}
}

func (l *DefaultLogger) Fatal(ctx context.Context, err error, format string, args ...any) {
	l.logger.ErrorContext(ctx, fmt.Sprintf("%s (%s)", err.Error(), fmt.Sprintf(format, args...)))
	panic(format)
}

func (l *DefaultLogger) With(ctx context.Context, key string, val any) context.Context {
	return context.WithValue(ctx, key, val)
}
