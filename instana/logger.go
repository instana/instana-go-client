package instana

import (
	"fmt"
	"log"
	"os"
)

// Logger is an interface for logging within the Instana client.
// It is compatible with popular logging libraries like logrus and zap.
type Logger interface {
	// Debug logs a debug message with optional key-value pairs
	Debug(msg string, keysAndValues ...interface{})

	// Info logs an info message with optional key-value pairs
	Info(msg string, keysAndValues ...interface{})

	// Warn logs a warning message with optional key-value pairs
	Warn(msg string, keysAndValues ...interface{})

	// Error logs an error message with optional key-value pairs
	Error(msg string, keysAndValues ...interface{})
}

// ClientLogLevel represents the logging level for the client
type ClientLogLevel int

const (
	// ClientLogLevelDebug enables debug logging
	ClientLogLevelDebug ClientLogLevel = iota
	// ClientLogLevelInfo enables info logging
	ClientLogLevelInfo
	// ClientLogLevelWarn enables warning logging
	ClientLogLevelWarn
	// ClientLogLevelError enables error logging
	ClientLogLevelError
	// ClientLogLevelNone disables all logging
	ClientLogLevelNone
)

// String returns the string representation of the log level
func (l ClientLogLevel) String() string {
	switch l {
	case ClientLogLevelDebug:
		return "DEBUG"
	case ClientLogLevelInfo:
		return "INFO"
	case ClientLogLevelWarn:
		return "WARN"
	case ClientLogLevelError:
		return "ERROR"
	case ClientLogLevelNone:
		return "NONE"
	default:
		return "UNKNOWN"
	}
}

// DefaultLogger is a simple logger implementation using the standard log package.
type DefaultLogger struct {
	logger   *log.Logger
	level    ClientLogLevel
	redacted []string // List of strings to redact from logs (e.g., API tokens)
}

// NewDefaultLogger creates a new DefaultLogger with the specified log level.
func NewDefaultLogger(level ClientLogLevel) *DefaultLogger {
	return &DefaultLogger{
		logger:   log.New(os.Stderr, "[instana-go-client] ", log.LstdFlags),
		level:    level,
		redacted: []string{},
	}
}

// SetRedactedStrings sets strings that should be redacted from log output
func (l *DefaultLogger) SetRedactedStrings(strings []string) {
	l.redacted = strings
}

// redact replaces sensitive strings in the message
func (l *DefaultLogger) redact(msg string) string {
	result := msg
	for _, s := range l.redacted {
		if s != "" && len(s) > 4 {
			// Replace with first 4 chars + asterisks
			result = replaceAll(result, s, s[:4]+"****")
		}
	}
	return result
}

// replaceAll is a simple string replacement function
func replaceAll(s, old, new string) string {
	result := ""
	for {
		i := indexOf(s, old)
		if i == -1 {
			result += s
			break
		}
		result += s[:i] + new
		s = s[i+len(old):]
	}
	return result
}

// indexOf finds the index of substring in string
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// formatMessage formats the message with key-value pairs
func (l *DefaultLogger) formatMessage(level ClientLogLevel, msg string, keysAndValues ...interface{}) string {
	formatted := fmt.Sprintf("[%s] %s", level.String(), l.redact(msg))

	if len(keysAndValues) > 0 {
		formatted += " |"
		for i := 0; i < len(keysAndValues); i += 2 {
			if i+1 < len(keysAndValues) {
				formatted += fmt.Sprintf(" %v=%v", keysAndValues[i], l.redact(fmt.Sprintf("%v", keysAndValues[i+1])))
			} else {
				formatted += fmt.Sprintf(" %v=<missing>", keysAndValues[i])
			}
		}
	}

	return formatted
}

// Debug logs a debug message
func (l *DefaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	if l.level <= ClientLogLevelDebug {
		l.logger.Println(l.formatMessage(ClientLogLevelDebug, msg, keysAndValues...))
	}
}

// Info logs an info message
func (l *DefaultLogger) Info(msg string, keysAndValues ...interface{}) {
	if l.level <= ClientLogLevelInfo {
		l.logger.Println(l.formatMessage(ClientLogLevelInfo, msg, keysAndValues...))
	}
}

// Warn logs a warning message
func (l *DefaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	if l.level <= ClientLogLevelWarn {
		l.logger.Println(l.formatMessage(ClientLogLevelWarn, msg, keysAndValues...))
	}
}

// Error logs an error message
func (l *DefaultLogger) Error(msg string, keysAndValues ...interface{}) {
	if l.level <= ClientLogLevelError {
		l.logger.Println(l.formatMessage(ClientLogLevelError, msg, keysAndValues...))
	}
}

// NoOpLogger is a logger that does nothing. Useful for disabling logging.
type NoOpLogger struct{}

// NewNoOpLogger creates a new NoOpLogger
func NewNoOpLogger() *NoOpLogger {
	return &NoOpLogger{}
}

// Debug does nothing
func (l *NoOpLogger) Debug(msg string, keysAndValues ...interface{}) {}

// Info does nothing
func (l *NoOpLogger) Info(msg string, keysAndValues ...interface{}) {}

// Warn does nothing
func (l *NoOpLogger) Warn(msg string, keysAndValues ...interface{}) {}

// Error does nothing
func (l *NoOpLogger) Error(msg string, keysAndValues ...interface{}) {}

// Made with Bob
