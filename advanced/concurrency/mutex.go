package concurrency

import (
	"fmt"
	"sync"
)

func MutexInGo() {
	fmt.Println("-- Mutex In Go --")
	counter := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for range 5 {
		// 5 goroutines are initialized here to update same counter variable
		// we can use wg.Go wrapper to reduce boilerplate code to add and reduce goroutines using wg
		// to be precise no need to write wg.Add() and wg.Done()
		wg.Go(func() {
			mutex.Lock()         // Locking the goroutine with Mutex
			defer mutex.Unlock() // Unlock at the end of the goroutine
			counter++
		})
	}
	wg.Wait()
	// the main goroutine completes its execution until it hits wg.Wait(), if we try to print
	// counter above wg.Wait() it will not work, why because before the scheduled goroutines try to update
	// the counter the print will already be completed.
	fmt.Println("Updated counter with Mutex", counter)
}
