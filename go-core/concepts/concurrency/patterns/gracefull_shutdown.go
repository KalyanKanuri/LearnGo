package concpatterns

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workJobs(ctx context.Context, wg *sync.WaitGroup, id int, jobs <-chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Worker %d processing %d\n", id, job)
			time.Sleep(time.Second)
		}
	}
}

func ExecGSDown() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)

	for id := range 2 {
		wg.Add(1)
		go workJobs(ctx, &wg, id, jobs)
	}

	go func() {
		for job := range 100 {
			jobs <- job + 1
		}
		close(jobs)
	}()

	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("All workers stopped gracefully")
}
