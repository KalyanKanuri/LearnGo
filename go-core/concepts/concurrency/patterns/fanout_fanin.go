package concpatterns

import (
	"fmt"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker ID %d, working on job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		res := job * job
		results <- res
	}
}

func ExecFOFI() {
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup

	for id := range 2 {
		wg.Add(1)
		go Worker(&wg, id+1, jobs, results)
	}

	go func() {
		for job := range 6 {
			jobN := job + 1
			jobs <- jobN
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
