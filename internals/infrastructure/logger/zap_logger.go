package logger

import (
	"os"
	"strings"

	"auptex.com/botnova/internals/application/ports"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	log *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {

	rotatingLog := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,   // MB before rotation
		MaxBackups: 5,    // number of old files to keep
		MaxAge:     30,   // days to keep old files
		Compress:   true, // gzip old files
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		CallerKey:      "caller",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(rotatingLog),
		zapcore.DebugLevel,
	)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	core := zapcore.NewTee(fileCore, consoleCore)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	return &ZapLogger{log: logger}, nil
}

func (z *ZapLogger) GetZapLogger() *zap.Logger {
	return z.log
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func (l *ZapLogger) Debug(msg string, fields ...ports.Field) {
	l.log.Debug(msg, toZapFields(fields)...)
}

func (l *ZapLogger) Info(msg string, fields ...ports.Field) {
	l.log.Info(msg, toZapFields(fields)...)
}

func (l *ZapLogger) Warn(msg string, fields ...ports.Field) {
	l.log.Warn(msg, toZapFields(fields)...)
}

func (l *ZapLogger) Error(msg string, fields ...ports.Field) {
	l.log.Error(msg, toZapFields(fields)...)
}

func (l *ZapLogger) With(fields ...ports.Field) ports.Logger {
	return &ZapLogger{
		log: l.log.With(toZapFields(fields)...),
	}
}

func (l *ZapLogger) Sync() error {
	return l.log.Sync()
}

func toZapFields(fields []ports.Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
