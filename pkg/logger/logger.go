package logger

import (
	"fmt"
	"log"
	"os"
)

// LogLevel represents the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Logger interface defines logging methods
type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

// SimpleLogger implements a simple logger
type SimpleLogger struct {
	level  LogLevel
	logger *log.Logger
}

// New creates a new logger instance
func New(level string) Logger {
	logLevel := parseLogLevel(level)
	return &SimpleLogger{
		level:  logLevel,
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// parseLogLevel parses string log level to LogLevel
func parseLogLevel(level string) LogLevel {
	switch level {
	case "debug", "DEBUG":
		return DEBUG
	case "info", "INFO":
		return INFO
	case "warn", "WARN":
		return WARN
	case "error", "ERROR":
		return ERROR
	default:
		return INFO
	}
}

// Debug logs a debug message
func (l *SimpleLogger) Debug(msg string, keysAndValues ...interface{}) {
	if l.level <= DEBUG {
		l.logWithLevel("DEBUG", msg, keysAndValues...)
	}
}

// Info logs an info message
func (l *SimpleLogger) Info(msg string, keysAndValues ...interface{}) {
	if l.level <= INFO {
		l.logWithLevel("INFO", msg, keysAndValues...)
	}
}

// Warn logs a warning message
func (l *SimpleLogger) Warn(msg string, keysAndValues ...interface{}) {
	if l.level <= WARN {
		l.logWithLevel("WARN", msg, keysAndValues...)
	}
}

// Error logs an error message
func (l *SimpleLogger) Error(msg string, keysAndValues ...interface{}) {
	if l.level <= ERROR {
		l.logWithLevel("ERROR", msg, keysAndValues...)
	}
}

// logWithLevel logs a message with the specified level
func (l *SimpleLogger) logWithLevel(level, msg string, keysAndValues ...interface{}) {
	logMsg := "[" + level + "] " + msg

	// Add key-value pairs if provided
	if len(keysAndValues) > 0 {
		logMsg += " |"
		for i := 0; i < len(keysAndValues); i += 2 {
			if i+1 < len(keysAndValues) {
				logMsg += " " + keysAndValues[i].(string) + "=" + formatValue(keysAndValues[i+1])
			}
		}
	}

	l.logger.Println(logMsg)
}

// formatValue formats a value for logging
func formatValue(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		return fmt.Sprint(v)
	}
}
