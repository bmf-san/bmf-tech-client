package logger

import (
	"log/slog"
	"os"
)

// Logger represents the singular of logger.
type Logger struct {
	*slog.Logger
}

// NewLogger creates a logger.
func NewLogger(level int) *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(level),
	})
	logger := slog.New(handler)
	return &Logger{
		logger,
	}
}
