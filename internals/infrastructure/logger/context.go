package logger

import (
	"context"

	"auptex.com/botnova/internals/application/ports"
)

type contextKey string

const loggerKey = contextKey("logger")

func WithContext(ctx context.Context, log ports.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

func FromContext(ctx context.Context) ports.Logger {
	if ctx == nil {
		return nil
	}
	if log, ok := ctx.Value(loggerKey).(ports.Logger); ok {
		return log
	}
	return nil
}
