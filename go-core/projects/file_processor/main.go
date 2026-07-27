package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

type JobStruct struct {
	FileName string
}

type ResultStruct struct {
	FileName   string
	WordsCount int
}

func fileWorker(wg *sync.WaitGroup, jobsQue <-chan JobStruct, resultsQue chan<- ResultStruct) {
	defer wg.Done()

	for job := range jobsQue {
		fileContent, err := os.ReadFile(job.FileName)
		words := strings.Fields(string(fileContent))
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		resultsQue <- ResultStruct{
			FileName:   job.FileName,
			WordsCount: len(words),
		}
	}
}

func main() {
	// variable initialization
	var wg sync.WaitGroup
	jobsQue := make(chan JobStruct, 6)
	resultsQue := make(chan ResultStruct, 6)
	filesDir := "./files"

	files, err := os.ReadDir(filesDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// set the number of workers to the number of CPU cores
	numWorkers := runtime.NumCPU()

	// initialize worker pool
	for range numWorkers {
		wg.Add(1)

		go fileWorker(&wg, jobsQue, resultsQue)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		jobsQue <- JobStruct{FileName: filesDir + "/" + file.Name()}
	}
	close(jobsQue)

	// wait for all workers to finish and close the results channel
	go func() {
		wg.Wait()
		close(resultsQue)
	}()

	// process results
	for result := range resultsQue {
		fmt.Printf("File: %s, Words Count: %d\n", result.FileName, result.WordsCount)
	}
}
