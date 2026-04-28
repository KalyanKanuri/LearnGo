package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Level int

const (
	Debug Level = iota
	Info
	Warning
	Error
	Critical
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARN"
	case Error:
		return "ERROR"
	case Critical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

type LogConfig struct {
	Out      io.Writer
	LogLevel Level
}

func NewLogger(out io.Writer, level Level) *LogConfig {
	if out == nil {
		out = os.Stdout
	}
	return &LogConfig{
		Out:      out,
		LogLevel: level,
	}
}

func (l *LogConfig) log(level Level, msg string, args ...any) {
	if level < l.LogLevel {
		return
	}
	ts := time.Now().Format(time.RFC3339)
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(l.Out, "%s\t%s\t%s\n", ts, level.String(), msg)
}

func (l *LogConfig) Debug(msg string, args ...any)    { l.log(Debug, msg, args...) }
func (l *LogConfig) Info(msg string, args ...any)     { l.log(Info, msg, args...) }
func (l *LogConfig) Warning(msg string, args ...any)  { l.log(Warning, msg, args...) }
func (l *LogConfig) Error(msg string, args ...any)    { l.log(Error, msg, args...) }
func (l *LogConfig) Critical(msg string, args ...any) { l.log(Critical, msg, args...) }
