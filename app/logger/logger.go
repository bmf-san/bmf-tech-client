package logger

import (
	"context"
	"log/slog"
	"os"
)

// Logger represents the singular of logger.
type Logger struct {
	*slog.Logger
}

// NewLogger creates a logger.
func NewLogger(level int) *Logger {
	handler := TraceIDHandler{slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(level),
	})}
	logger := slog.New(handler)
	return &Logger{
		logger,
	}
}

// WithTraceID returns a context with a trace id.
func (l *Logger) WithTraceID(ctx context.Context, tid string) context.Context {
	return context.WithValue(ctx, ctxTraceIDKey, tid)
}

// TraceIDHandler represents the singular of trace id handler.
type TraceIDHandler struct {
	slog.Handler
}

type ctxTraceID struct{}

var ctxTraceIDKey = ctxTraceID{}

// Handle implements slog.Handler.
func (t TraceIDHandler) Handle(ctx context.Context, r slog.Record) error {
	tid, ok := ctx.Value(ctxTraceIDKey).(string)
	if ok {
		r.AddAttrs(slog.String("trace_id", tid))
	}
	return t.Handler.Handle(ctx, r)
}
