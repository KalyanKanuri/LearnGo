package concurrency

import (
	"fmt"
	"time"
)

func executeTask(msg string, delay time.Duration) {
	fmt.Printf("Task execution start time %s\n", time.Now())
	fmt.Printf("sleeping for %s time\n", delay)
	time.Sleep(delay)
	fmt.Printf("Task message: %s\n", msg)
	fmt.Printf("Task execution end time %s\n", time.Now())
}

func GoRoutinesInGo() {
	fmt.Println("-- Go Routines in Go --")
	go executeTask("Executing demo go routine", time.Second)
	go executeTask("Executing demo go routine", 4*time.Second)
	go executeTask("Executing demo go routine", 4*time.Second)
	go executeTask("Executing demo go routine", 4*time.Second)
	// we have to give sufficient time for go routine to complete
	time.Sleep(2 * time.Second) // wait for go routine to complete
}
