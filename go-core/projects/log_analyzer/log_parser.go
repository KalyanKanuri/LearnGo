package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseLogLines(file io.Reader) (LogSummary, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	logSummary := LogSummary{
		TotalLines: lineCount,
		INFO:       0,
		WARN:       0,
		ERROR:      0,
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		logFields := strings.Fields(line)
		level := logFields[2]

		if len(logFields) < 3 {
			fmt.Printf("Skipping malformed line at %d: %s", lineCount, line)
			continue
		}

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
