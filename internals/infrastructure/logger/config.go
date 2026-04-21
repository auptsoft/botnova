package logger

type Config struct {
	Environment     string // e.g., "development", "production"
	ServiceName     string
	ConsoleLogLevel string // e.g., "debug", "info", "warn", "error"
	FileLogLevel    string
	Filename        string
}
