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

func main() {
	file, err := os.Open("logfile.log")
	if err != nil {
		fmt.Println("Error while opening file", err)
		return
	}
	defer file.Close()

	logSummary, err := ParseLogLines(file)
	if err != nil {
		fmt.Println("Error Parsing log file", err)
	}

	fmt.Println("============ Log Summary ============")
	fmt.Printf(
		"\nTotal Lines: %d\n\nInfo:%d\nWarn:%d\nError:%d\n\n",
		logSummary.TotalLines, logSummary.INFO, logSummary.WARN, logSummary.ERROR,
	)
	fmt.Println("======================================")
}
