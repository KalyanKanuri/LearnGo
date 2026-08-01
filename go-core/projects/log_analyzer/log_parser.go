package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseLogLines(file io.Reader) map[string]int {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	logData := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		logFields := strings.Fields(line)

		if len(logFields) < 3 {
			fmt.Printf("Skipping malformed line at %d", lineCount)
			continue
		}

		logData[logFields[2]]++
	}
	logData["totalLines"] = lineCount

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file", err)
		return map[string]int{}
	}
	return logData
}
