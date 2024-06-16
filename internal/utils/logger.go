package utils

import (
	"log/slog"

	"github.com/getsentry/sentry-go"
)

// CustomLogger wraps slog and sends messages to Sentry based on the log level.
type CustomLogger struct {
	*slog.Logger
}

// NewCustomLogger creates a new instance of CustomLogger.
func NewCustomLogger(logger *slog.Logger) *CustomLogger {
	return &CustomLogger{Logger: logger}
}

// Error logs an error message with slog and captures it with Sentry.
func (l *CustomLogger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...) // Log with slog
	sentry.CaptureMessage(msg)   // Send to Sentry
}

var Logger *CustomLogger

func InitLogger() {
	Logger = NewCustomLogger(slog.Default())
}
