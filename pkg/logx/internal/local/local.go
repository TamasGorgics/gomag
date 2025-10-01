package local

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
)

const (
	ColorReset = "\033[0m"
	ColorWhite = "\033[37m"
	ColorGray  = "\033[90m"

	ColorBrightRed    = "\033[91m"
	ColorBrightGreen  = "\033[92m"
	ColorBrightYellow = "\033[93m"
	ColorBrightBlue   = "\033[94m"
	ColorBrightCyan   = "\033[96m"
)

// LocalLogger is a human-readable logger for local development
type LocalLogger struct {
	level  slog.Level
	logger *slog.Logger
}

// NewLocalLogger creates a new local development logger with pretty formatting
func NewLocalLogger(level slog.Level) *LocalLogger {
	handler := &LocalHandler{
		level: level,
	}

	return &LocalLogger{
		level:  level,
		logger: slog.New(handler),
	}
}

// LocalHandler implements slog.Handler for pretty local development logging
type LocalHandler struct {
	level slog.Level
}

func (h *LocalHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *LocalHandler) Handle(ctx context.Context, r slog.Record) error {
	// Get caller information
	pc, file, line, ok := runtime.Caller(4)
	caller := ""
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			// Extract just the function name
			fullName := fn.Name()
			parts := strings.Split(fullName, ".")
			caller = parts[len(parts)-1]
		}
		// Extract just the filename
		fileParts := strings.Split(file, "/")
		file = fileParts[len(fileParts)-1]
	}

	// Format timestamp
	timestamp := r.Time.Format("15:04:05.000")

	// Get level color
	levelColor := h.getLevelColor(r.Level)

	// Format the message
	message := r.Message

	// Format attributes
	var attrs []string
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, fmt.Sprintf("%s=%v", a.Key, a.Value))
		return true
	})

	// Build the log line
	var parts []string

	// Timestamp
	parts = append(parts, fmt.Sprintf("%s%s%s", ColorGray, timestamp, ColorReset))

	// Level with color
	levelStr := fmt.Sprintf("%s%s%s", levelColor, r.Level.String(), ColorReset)
	parts = append(parts, levelStr)

	// Caller info
	if caller != "" {
		callerInfo := fmt.Sprintf("%s%s:%d %s()%s", ColorBrightCyan, file, line, caller, ColorReset)
		parts = append(parts, callerInfo)
	}

	// Message
	parts = append(parts, fmt.Sprintf("%s%s%s", ColorWhite, message, ColorReset))

	// Attributes
	if len(attrs) > 0 {
		attrStr := fmt.Sprintf("%s[%s]%s", ColorGray, strings.Join(attrs, ", "), ColorReset)
		parts = append(parts, attrStr)
	}

	// Print the formatted log
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(parts, " "))

	return nil
}

func (h *LocalHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// For simplicity, we'll just return the same handler
	// In a more complex implementation, you might want to store these attrs
	return h
}

func (h *LocalHandler) WithGroup(name string) slog.Handler {
	// For simplicity, we'll just return the same handler
	// In a more complex implementation, you might want to handle groups
	return h
}

func (h *LocalHandler) getLevelColor(level slog.Level) string {
	switch {
	case level >= slog.LevelError:
		return ColorBrightRed
	case level >= slog.LevelWarn:
		return ColorBrightYellow
	case level >= slog.LevelInfo:
		return ColorBrightGreen
	case level >= slog.LevelDebug:
		return ColorBrightBlue
	default:
		return ColorGray
	}
}

// Logger interface implementation
func (l *LocalLogger) Debug(ctx context.Context, format string, args ...any) {
	l.logger.DebugContext(ctx, format, args...)
}

func (l *LocalLogger) Info(ctx context.Context, format string, args ...any) {
	l.logger.InfoContext(ctx, format, args...)
}

func (l *LocalLogger) Warn(ctx context.Context, format string, args ...any) {
	l.logger.WarnContext(ctx, format, args...)
}

func (l *LocalLogger) Error(ctx context.Context, err error, format string, args ...any) {
	if err != nil {
		l.logger.ErrorContext(ctx, fmt.Sprintf("%s (%s)", err.Error(), fmt.Sprintf(format, args...)))
	} else {
		l.logger.ErrorContext(ctx, format, args...)
	}
}

func (l *LocalLogger) Fatal(ctx context.Context, err error, format string, args ...any) {
	l.logger.ErrorContext(ctx, fmt.Sprintf("%s (%s)", err.Error(), fmt.Sprintf(format, args...)))
	panic(fmt.Sprintf(format, args...))
}

func (l *LocalLogger) With(ctx context.Context, key string, val any) context.Context {
	return context.WithValue(ctx, key, val)
}
