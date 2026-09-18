package middleware

import (
	"context"
	"net/http"
	"time"
)

type Logger interface {
	Info(ctx context.Context, msg string, args ...any)
}

func Logging(logger Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &writeWrapper{ResponseWriter: w}
		next.ServeHTTP(ww, r)
		logger.Info(r.Context(), "Request completed", "method", r.Method, "url", r.URL.String(), "requestID", r.Context().Value("requestID"), "statusCode", ww.statusCode, "duration_ns", time.Since(start).Nanoseconds())
	})
}

type writeWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *writeWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
