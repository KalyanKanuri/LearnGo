package concurrency

import (
	"fmt"
	"sync"
)

// this is the most common declaration used in go community for waitgroups
var wg sync.WaitGroup

func runTask(iter int, wg *sync.WaitGroup) {
	// we have to signal that the goroutine execution is completed using below
	// if not the goroutine might stay alive
	defer wg.Done()
	fmt.Println("Executing iteration", iter)
}

func WaitGroupsInGo() {
	fmt.Println("-- WaitGroups In Go --")
	for i := range 5 {
		wg.Add(1)
		go runTask(i, &wg)
	}
	// even though if we wait here it's useless we have to use wait in main goroutine
	// just for execution purpose & understandability we have written to wait here
	wg.Wait() // without wait this goroutine will not wait for other goroutines to complete
}
