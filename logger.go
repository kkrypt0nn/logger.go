package logging

import (
	"fmt"
	"github.com/kkrypt0nn/logger.go/terminal"
	"os"
)

var (
	CurrentLoggingLevel   = NONE
	CurrentDateFormat     = "Jan 02, 2006"
	CurrentDatetimeFormat = "Jan 02, 2006 15:04:05"
	CurrentTimeFormat     = "15:04:05"
)

// Logger is represents a logger structure.
type Logger struct {
	// Styling describes whether the logger should style the logged message.
	Styling bool
	// Prefix is the prefix before the logged message.
	Prefix string
	// LogFile is the log file to write the messages into.
	LogFile *os.File
}

// NewLogger creates a new logger.
func NewLogger() *Logger {
	return &Logger{
		Styling: true,
		Prefix:  "${datetime} ${level:color}${level:name}${reset}: ",
		LogFile: nil,
	}
}

// SetLogFile will set the log file to write logs into.
func (l *Logger) SetLogFile(file *os.File) {
	l.LogFile = file
}

// SetDateFormat sets the format to use when logging the date.
func (l *Logger) SetDateFormat(format string) {
	CurrentDateFormat = format
}

// SetDatetimeFormat sets the format to use when logging the date and time.
func (l *Logger) SetDatetimeFormat(format string) {
	CurrentDatetimeFormat = format
}

// SetTimeFormat sets the format to use when logging the time.
func (l *Logger) SetTimeFormat(format string) {
	CurrentTimeFormat = format
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
	CurrentLoggingLevel = level
	// So that Print and Println do not keep the previous logging level.
	defer l.SetLoggingLevel(NONE)
	message = l.Prefix + message + terminal.RESET
	message = AddVariables(message)
	if l.Styling && terminal.AreColorsSupported() {
		message = AddStyling(message)
	}
	fmt.Println(message)
	if l.LogFile != nil {
		message = RemoveStyling(message)
		_, err := l.LogFile.WriteString(message + "\n")
		if err != nil {
			panic(err)
		}
	}
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

// Log will log with the given level.
func (l *Logger) Log(level Level, message string) {
	l.doLog(level, message)
}

// SetLoggingLevel sets the logging level.
func (l *Logger) SetLoggingLevel(level Level) {
	CurrentLoggingLevel = level
}

// Print simply prints the message, without logging level.
func (l *Logger) Print(message string) {
	message = AddVariables(message)
	if l.Styling && terminal.AreColorsSupported() {
		message = AddStyling(message)
	}
	fmt.Print(message + terminal.RESET)
}

// Println simply prints the message with a new line, without logging level.
func (l *Logger) Println(message string) {
	message = AddVariables(message)
	if l.Styling && terminal.AreColorsSupported() {
		message = AddStyling(message)
	}
	fmt.Println(message + terminal.RESET)
}
