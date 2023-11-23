package glair

import (
	"fmt"
	"os"
)

// LogLevel represents logging level of GLAIR Vision SDK
type LogLevel int

const (
	// LevelNone sets the logger to disables logging entirely
	LevelNone LogLevel = 0
	// LevelError sets the logger to show only error messages
	LevelError LogLevel = 1
	// LevelWarn sets the logger to show warning and error messages
	LevelWarn LogLevel = 2
	// LevelInfo sets the logger to show information, warning, and error messages
	LevelInfo LogLevel = 3
	// LevelDebug sets the logger to display any messages, including debugging statements
	LevelDebug LogLevel = 4
)

// Logger represents base contract for logging functionalities
type Logger interface {
	// Debugf prints debug message with the provided format to stdout
	Debugf(format string, val ...interface{})
	// Infof prints informational message with the provided format to stdout
	Infof(format string, val ...interface{})
	// Warnf prints warning message with the provided format to stdout
	Warnf(format string, val ...interface{})
	// Errorf prints error message with the provided format to stderr
	Errorf(format string, val ...interface{})
}

// LeveledLogger represents logger instance that will be used by GLAIR Vision
// Go SDK
type LeveledLogger struct {
	Level LogLevel
}

func (l *LeveledLogger) Debugf(format string, val ...interface{}) {
	if l.Level >= LevelDebug {
		fmt.Printf("[DEBUG] "+format, val)
	}
}

func (l *LeveledLogger) Infof(format string, val ...interface{}) {
	if l.Level >= LevelInfo {
		fmt.Printf("[INFO] "+format, val)
	}
}

func (l *LeveledLogger) Warnf(format string, val ...interface{}) {
	if l.Level >= LevelWarn {
		fmt.Printf("[WARN] "+format, val)
	}
}

func (l *LeveledLogger) Errorf(format string, val ...interface{}) {
	if l.Level >= LevelError {
		fmt.Fprintf(os.Stderr, "[ERROR] "+format, val)
	}
}
