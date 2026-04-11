package core

import (
	"context"
	"fmt"
	"time"
)

func ContextInGo() {
	fmt.Println("\n*******  Context in Go ******* ")
	fmt.Println("-- Context with timeout simulation --")
	parentCtx := context.Background()

	timeCtx, timeCan := context.WithTimeout(parentCtx, 1*time.Second)
	defer timeCan()

	timeoutDoneChan := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(timeoutDoneChan)
	}()

	select {
	case <-timeCtx.Done():
		fmt.Println("context time out simulation completed - err:", timeCtx.Err())
	case <-timeoutDoneChan:
		fmt.Println("context process simulation completed")
	}

	fmt.Println("-- Context with Cancel Simulation --")
	cancelCtx, cancCan := context.WithCancel(parentCtx)

	cancelDoneChan := make(chan struct{})

	go func(ctx context.Context) {
		defer close(cancelDoneChan)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancel simulation completed - err:", ctx.Err())
				return
			default:
				fmt.Println("default stmnt")
			}
		}
	}(cancelCtx)

	time.Sleep(time.Millisecond)
	cancCan() // manually cancel

	<-cancelDoneChan
	fmt.Println("cancel ctx done")

	fmt.Println("-- Context with Deadline and value Simulation --")
	type key string
	const userID key = "user_id"

	valCtx := context.WithValue(parentCtx, userID, "usr-123")
	deadline := time.Now().Add(3 * time.Second)
	dlValCtx, deadlineCan := context.WithDeadline(valCtx, deadline)
	defer deadlineCan()

	deadlineDoneChan := make(chan struct{})

	go func(ctx context.Context) {
		defer close(deadlineDoneChan)

		select {
		case <-ctx.Done():
			fmt.Println("Context deadline simulation completed - err:", ctx.Err(), ctx.Value(userID))
		case <-time.After(time.Second):
			fmt.Println("Context deadline process completed", ctx.Value(userID))
		}
	}(dlValCtx)

	<-deadlineDoneChan
}
