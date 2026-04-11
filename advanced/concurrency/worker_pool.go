package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Worker Pool pattern
 * -------------------
 * * A channel for queueing jobs or tasks
 * * A channel for collecting results [optional]
 * * A worker or processor to process the task
 * this combination is called worker pool we can control the number of goroutines
 * getting created and reuse them for the work instead of spanning a new one for each task.
 */
func worker(wg *sync.WaitGroup, jobs <-chan int, results chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Working on Job", job)
		time.Sleep(time.Second)
		fmt.Println("Worker Completed", job)
		results <- job * 2 // just to show some change happened we are multiplying this
	}
}

func workPoolExec() {
	jobs := make(chan int)
	results := make(chan int)

	maxWorkers := 5
	for i := 0; i <= maxWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}

	go func() {
		defer close(jobs)
		for i := range maxWorkers {
			jobs <- i
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func WorkerPoolInGo() {
	fmt.Println("-- Worker Pool pattern in Go --")
	workPoolExec()
}
