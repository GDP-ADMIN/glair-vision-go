package glair

import (
	"testing"
)

// TestLeveledLogger_AllLevels exercises each logging method at LevelDebug
// (all branches active) and LevelNone (all branches inactive).
func TestLeveledLogger_AllLevels(t *testing.T) {
	loggers := []*LeveledLogger{
		{Level: LevelDebug},
		{Level: LevelNone},
	}

	for _, l := range loggers {
		l.Debugf("debug %s", "msg")
		l.Infof("info %s", "msg")
		l.Warnf("warn %s", "msg")
		l.Errorf("error %s", "msg")
	}
}

// TestLeveledLogger_IntermediateLevels exercises intermediate log levels to
// cover the remaining branches (LevelError, LevelWarn, LevelInfo).
func TestLeveledLogger_IntermediateLevels(t *testing.T) {
	for _, lvl := range []LogLevel{LevelError, LevelWarn, LevelInfo} {
		l := &LeveledLogger{Level: lvl}
		l.Debugf("debug %s", "msg")
		l.Infof("info %s", "msg")
		l.Warnf("warn %s", "msg")
		l.Errorf("error %s", "msg")
	}
}
