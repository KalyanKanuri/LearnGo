// Package logger provides a simple logging utility using the zerolog library.
package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New initializes and returns a new zerolog.Logger instance configured to log to the console with timestamps in RFC3339 format. The logger's output is directed to stdout in release mode and stderr otherwise.
func New() zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	if os.Getenv("GIN_MODE") == "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	log.Logger = log.Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	)
	return log.Logger
}
