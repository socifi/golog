package log

import (
	"bytes"
	"errors"
	"strings"
)

// ErrInvalidLevel is returned if the severity level is invalid.
var ErrInvalidLevel = errors.New("invalid level")

// Level of severity.
type Level int

// Log levels.
const (
	InvalidLevel Level = -1
	DebugLevel = 100
	InfoLevel = 200
	NoticeLevel = 250
	WarnLevel = 300
	ErrorLevel = 400
	CriticalLevel = 500
	AlertLevel = 550
	FatalLevel = 600
	EmergencyLevel = 600
)

var levelNames = [...]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	NoticeLevel: "notice",
	WarnLevel:  "warn",
	ErrorLevel: "error",
	CriticalLevel: "critical",
	AlertLevel: "alert",
	FatalLevel: "emergency",
}

var levelStrings = map[string]Level{
	"debug":   DebugLevel,
	"info":    InfoLevel,
	"notice":  NoticeLevel,
	"warn":    WarnLevel,
	"warning": WarnLevel,
	"error":   ErrorLevel,
	"critical":CriticalLevel,
	"alert":   AlertLevel,
	"emergency":FatalLevel,
}

// String implementation.
func (l Level) String() string {
	return levelNames[l]
}

// String implementation.
func (l Level) Int() int {
	return int(l)
}

// MarshalJSON implementation.
func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.String() + `"`), nil
}

// UnmarshalJSON implementation.
func (l *Level) UnmarshalJSON(b []byte) error {
	v, err := ParseLevel(string(bytes.Trim(b, `"`)))
	if err != nil {
		return err
	}

	*l = v
	return nil
}

// ParseLevel parses level string.
func ParseLevel(s string) (Level, error) {
	l, ok := levelStrings[strings.ToLower(s)]
	if !ok {
		return InvalidLevel, ErrInvalidLevel
	}

	return l, nil
}

// MustParseLevel parses level string or panics.
func MustParseLevel(s string) Level {
	l, err := ParseLevel(s)
	if err != nil {
		panic("invalid log level")
	}

	return l
}
