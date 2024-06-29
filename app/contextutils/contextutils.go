package contextutils

import "context"

type ctxTraceID struct{}

var ctxTraceIDKey = ctxTraceID{}

// SetTraceID sets a trace ID to a context.
func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, ctxTraceIDKey, traceID)
}

// GetTraceID returns a trace ID from a context.
func GetTraceID(ctx context.Context) string {
	if v, ok := ctx.Value(ctxTraceIDKey).(string); ok {
		return v
	}
	return ""
}
