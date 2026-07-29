package main

import (
	"context"
	"net/http"
	"time"
)

func HealthChecker(ctx context.Context, url string) Result {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Result{
			URL:        url,
			StatusCode: 0,
			Latency:    0,
			Err:        err,
		}
	}

	startTime := time.Now()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return Result{
			URL:        url,
			StatusCode: 0,
			Latency:    time.Since(startTime),
			Err:        err,
		}
	}
	defer resp.Body.Close()

	return Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Latency:    time.Since(startTime),
		Err:        err,
	}
}
