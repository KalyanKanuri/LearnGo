package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

type Job struct {
	URL string
}

type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	Err        error
}

type ResultSummary struct {
	TotalURLs      int
	Healthy        int
	Failed         int
	AverageLatency time.Duration
	Slowest        string
	Fastest        string
}

func resultAggregator(results <-chan Result) ResultSummary {
	totalUrls := 0
	healthy := 0
	failed := 0
	slowest := ""
	fastest := ""
	maxLatency := 0
	var totalLatency time.Duration
	var avgLatency time.Duration
	minLatency := time.Duration(math.MaxInt64)

	for res := range results {
		totalUrls++
		if res.StatusCode == 200 {
			healthy++
		} else {
			failed++
		}
		totalLatency += res.Latency

		if res.Latency > time.Duration(maxLatency) {
			maxLatency = int(res.Latency)
			slowest = res.URL
		}

		if res.Latency < time.Duration(minLatency) {
			minLatency = res.Latency
			fastest = res.URL
		}
	}

	if totalUrls != 0 {
		avgLatency = totalLatency / time.Duration(totalUrls)
	}

	return ResultSummary{
		TotalURLs:      totalUrls,
		Healthy:        healthy,
		Failed:         failed,
		AverageLatency: avgLatency,
		Slowest:        slowest,
		Fastest:        fastest,
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)
	var wg sync.WaitGroup

	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	numWorkers := min(len(urls), 8)
	for range numWorkers {
		wg.Add(1)
		go Worker(ctx, &wg, jobs, results)
	}

	go func() {
		for _, url := range urls {
			jobs <- Job{
				URL: url,
			}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	summary := resultAggregator(results)
	fmt.Println("======= Summary =======")
	fmt.Printf(
		"Total URLs 		:%d\n"+
			"Healthy 		:%d\n"+
			"Failed 			:%d\n"+
			"Average Latency 	:%v\n"+
			"Slowest 		:%s\n"+
			"Fastest 		:%s\n",
		summary.TotalURLs, summary.Healthy,
		summary.Failed, summary.AverageLatency,
		summary.Slowest, summary.Fastest,
	)
}
