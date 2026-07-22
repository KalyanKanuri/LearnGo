package concurrency

import (
	"fmt"
	"time"
)

func ExecuteTimeCase() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine!"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No message received within 1 second.")
	}
}
