package concpatterns

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func ExecWP() {
	jobs := make(chan int)

	for i := range 2 {
		wg.Add(1)
		go worker(i+1, jobs)
	}

	for job := range 5 {
		jobs <- job + 1
	}

	close(jobs)
	wg.Wait()
}
