package main

import (
	"fmt"
	"os"
)

type LogSummary struct {
	TotalLines int
	INFO       int
	WARN       int
	ERROR      int
}

func buildSummary(logData map[string]int) LogSummary {
	logSummary := LogSummary{
		TotalLines: logData["totalLines"],
		INFO:       logData["INFO"],
		WARN:       logData["WARN"],
		ERROR:      logData["ERROR"],
	}
	return logSummary
}

func main() {
	file, err := os.Open("logfile.log")
	if err != nil {
		fmt.Println("Error while opening file", err)
		return
	}
	defer file.Close()

	logData := ParseLogLines(file)

	fmt.Println("============ Log Summary ============")
	logSummary := buildSummary(logData)
	fmt.Printf(
		"\nTotal Lines: %d\n\nInfo:%d\nWarn:%d\nError:%d\n\n",
		logSummary.TotalLines, logSummary.INFO, logSummary.WARN, logSummary.ERROR,
	)
	fmt.Println("======================================")
}
