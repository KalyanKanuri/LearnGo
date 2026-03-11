package glog

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// time package accepts 2006 and time pattern to render the time in given pattern
// any year other than 2006 does not work results year as 11003
const TIME_FORMAT = "2006-01-02 15:04:05"

func Log(level LogLevel, msg string) {
	loglevels := []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
	now := time.Now()
	fmt.Printf("%s	%s	%s\n", now.Format(TIME_FORMAT), loglevels[level], msg)
}
