package ports

import "go.uber.org/zap"

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)

	With(fields ...Field) Logger
	Sync() error

	GetZapLogger() *zap.Logger
}

type Field struct {
	Key   string
	Value any
}
