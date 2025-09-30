package manager

import "context"

type Logger interface {
	Info(ctx context.Context, format string, args ...any)
	Error(ctx context.Context, err error, format string, args ...any)
}
