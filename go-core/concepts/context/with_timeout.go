package contextGo

import (
	"context"
	"fmt"
	"time"
)

func ContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context timed out:", ctx.Err())
				return
			default:
				fmt.Println("Working with Timeout")
				time.Sleep(500*time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
