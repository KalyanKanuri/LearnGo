package contextGo

import (
	"context"
	"fmt"
	"time"
)

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker Stopped")
				return
			default:
				fmt.Println("Working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
