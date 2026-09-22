package timepkg

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func ExecTimer() {
	// Example usage of time.Timer
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	var wg sync.WaitGroup
	workEvents := make(chan int)

	go func() {
		defer close(workEvents)
		for i := range 10 {
			workEvents <- i + 1
		}
	}()

	wg.Go(func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("timer completed")
				return
			case e, ok := <-workEvents:
				if !ok {
					fmt.Println("All events processed")
					return
				}
				fmt.Println("event from worker", e)
			}
		}
	})

	fmt.Println("This will be printed even before or while the goroutine is getting executed")
	wg.Wait()
	fmt.Println("This will be printed once the goroutine is executed")
}

func ExecTicker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var wg sync.WaitGroup
	done := make(chan bool)

	wg.Go(func() {
		time.Sleep(5 * time.Second)
		done <- true
	})

	wg.Go(func() {
		for {
			select {
			case <-done:
				fmt.Println("Work Done")
				return
			case t := <-ticker.C:
				fmt.Printf("Executing task at %d:%d:%d\n", t.Hour(), t.Minute(), t.Second())
			}
		}

	})
	wg.Wait()
}

func ExecTimerTicker() {
	var wg sync.WaitGroup
	timer := time.NewTimer(10 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer timer.Stop()
	defer ticker.Stop()

	wg.Go(func() {
		for {
			select {
			case <-ticker.C:
				resp, err := http.Get("http://localhost:8080")
				if err != nil {
					fmt.Println("Error getting health check status", err)
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					if err == io.EOF {
						fmt.Printf("API Response: %s\n", string(body))
					}
					fmt.Println("Error reading response body: ", err)
					return
				}
				fmt.Printf("API Response: %s\n", string(body))
			case <-timer.C:
				fmt.Println("Health check Worker stopped after 10 seconds")
				return
			}
		}
	})

	wg.Wait()
}
