package main

import (
	"context"
	"sync"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			resp := HealthChecker(ctx, job.URL)

			result := Result{
				URL:        job.URL,
				StatusCode: resp.StatusCode,
				Err:        resp.Err,
			}
			results <- result
		}
	}
}
