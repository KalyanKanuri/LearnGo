package contextGo

import (
	"context"
	"fmt"
	"time"
)

func ContextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context Deadline reached:", ctx.Err())
				return
			default:
				fmt.Println("Working with Deadline")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
