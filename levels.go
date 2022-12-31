package logging

import "github.com/kkrypt0nn/logger.go/terminal"

// Level represents the level of logging - https://logging.apache.org/log4j/1.2/apidocs/org/apache/log4j/class-use/Level.html
type Level int

const (
	// DEBUG level designates fine-grained informational events that are most useful to debug an application.
	DEBUG Level = iota
	// INFO level designates informational messages that highlight the progress of the application at coarse-grained level.
	INFO
	// WARN level designates potentially harmful situations.
	WARN
	// ERROR level designates error events that might still allow the application to continue running.
	ERROR
	// FATAL level designates very severe error events that will presumably lead the application to abort.
	FATAL
	// TRACE level designates finer-grained informational events than the DEBUG.
	TRACE
)

var (
	// LevelColors convert a logging level to its color.
	LevelColors = map[Level]string{
		DEBUG: terminal.WHITE + terminal.BACK_BLACK,
		INFO:  terminal.WHITE + terminal.BACK_BLUE,
		WARN:  terminal.WHITE + terminal.BACK_YELLOW,
		ERROR: terminal.WHITE + terminal.BACK_RED,
		FATAL: terminal.WHITE + terminal.BOLD + terminal.BACK_RED,
		TRACE: terminal.BLACK + terminal.BACK_DARKGRAY,
	}
	// LevelNames convert a logging level to its name.
	LevelNames = map[Level]string{
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
		FATAL: "FATAL",
		TRACE: "TRACE",
	}
	// ShortLevelNames convert a logging level to its short name.
	ShortLevelNames = map[Level]string{
		DEBUG: "dbg",
		INFO:  "inf",
		WARN:  "war",
		ERROR: "err",
		FATAL: "fat",
		TRACE: "tra",
	}
)

// GetLevelColor returns the color of the level.
func GetLevelColor(id Level) string {
	if !terminal.AreColorsSupported() {
		return ""
	}
	return LevelColors[id]
}

// GetLevelName returns the name of the level.
func GetLevelName(id Level) string {
	return LevelNames[id]
}

// GetLevelShortName returns the short name of the level.
func GetLevelShortName(id Level) string {
	return ShortLevelNames[id]
}
