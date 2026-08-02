package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseLogLines(file io.Reader) (LogSummary, error) {
	scanner := bufio.NewScanner(file)
	var logSummary LogSummary

	for scanner.Scan() {
		line := scanner.Text()
		logSummary.TotalLines++
		logFields := strings.Fields(line)

		if len(logFields) < 3 {
			fmt.Printf("Skipping malformed line at %d: %s", logSummary.TotalLines, line)
			continue
		}
		level := logFields[2]

		switch level {
		case "INFO":
			logSummary.INFO++
		case "WARN":
			logSummary.WARN++
		case "ERROR":
			logSummary.ERROR++
		}
	}

	if err := scanner.Err(); err != nil {
		return LogSummary{}, err
	}

	return logSummary, nil
}
