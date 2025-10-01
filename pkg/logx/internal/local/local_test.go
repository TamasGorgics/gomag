package local_test

import (
	"context"
	"errors"
	"testing"

	"github.com/TamasGorgics/gomag/pkg/logx"
)

func TestLocalLogger(t *testing.T) {
	// Initialize the local logger for development
	_ = logx.InitLocalLogger()

	ctx := context.Background()

	// Demonstrate different log levels with pretty formatting
	logx.Debug(ctx, "This is a debug message with some details")
	logx.Info(ctx, "Application started successfully")
	logx.Warn(ctx, "This is a warning message")

	// Demonstrate error logging
	err := errors.New("database connection failed")
	logx.Error(ctx, err, "Failed to connect to database")

	// Demonstrate context with additional data
	ctx = logx.With(ctx, "user_id", "12345")
	ctx = logx.With(ctx, "request_id", "req-abc-123")
	logx.Info(ctx, "Processing user request")

	// Test with structured logging attributes
	logx.Info(ctx, "User action completed", "action", "login", "duration_ms", 150)

	// Demonstrate different log levels in context
	logx.Debug(ctx, "Debug info: processing step 1")
	logx.Info(ctx, "User authentication successful")
	logx.Warn(ctx, "Rate limit approaching for user", "user_id", "12345")

	// This would normally cause a panic, but we'll comment it out for the example
	// logx.Fatal(ctx, err, "Critical system failure")

	logx.Info(ctx, "Example completed - check out the pretty formatted logs above!")
}
