package logging

import (
	"fmt"
	"github.com/kkrypt0nn/logger.go/terminal"
)

// Logger is represents a logger structure.
type Logger struct {
	// Styling describes whether the logger should style the logged message.
	Styling bool
	// Prefix is the prefix before the logged message.
	Prefix string
}

// NewLogger creates a new logger.
func NewLogger() *Logger {
	return &Logger{
		Styling: true,
		Prefix:  "${datetime} ${level:color}${level:name}${reset}: ",
	}
}

// SetStyling sets the styling setting of the logger.
func (l *Logger) SetStyling(styling bool) {
	l.Styling = styling
}

// SetPrefix sets the prefix before the log message.
func (l *Logger) SetPrefix(prefix string) {
	l.Prefix = prefix
}

func (l *Logger) doLog(level Level, message string) {
	message = l.Prefix + message + terminal.RESET
	message = AddVariables(message)
	if l.Styling && terminal.AreColorsSupported() {
		message = AddStyling(message)
	}
	message = AddLoggingLevel(message, level)
	fmt.Println(message)
}

// Debug prints a debug message.
func (l *Logger) Debug(message string) {
	l.doLog(DEBUG, message)
}

// Info prints an info message.
func (l *Logger) Info(message string) {
	l.doLog(INFO, message)
}

// Warn prints a warning message.
func (l *Logger) Warn(message string) {
	l.doLog(WARN, message)
}

// Error prints an error message.
func (l *Logger) Error(message string) {
	l.doLog(ERROR, message)
}

// Fatal prints a fatal error message.
func (l *Logger) Fatal(message string) {
	l.doLog(FATAL, message)
}

// Trace prints a tracing message.
func (l *Logger) Trace(message string) {
	l.doLog(TRACE, message)
}

// Print simply prints the message.
func (l *Logger) Print(message string) {
	message = AddVariables(message)
	if l.Styling && terminal.AreColorsSupported() {
		message = AddStyling(message)
	}
	fmt.Print(message + terminal.RESET)
}

// Println simply prints the message with a new line.
func (l *Logger) Println(message string) {
	message = AddVariables(message)
	if l.Styling && terminal.AreColorsSupported() {
		message = AddStyling(message)
	}
	fmt.Println(message + terminal.RESET)
}
