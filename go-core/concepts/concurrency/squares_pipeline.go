package concurrency

import "fmt"

func GenerateSqrs() {
	numsCh := make(chan int)
	sqrsCh := make(chan int)

	go func() {
		for i := range 10 {
			numsCh <- i + 1
		}
		close(numsCh)
	}()

	go func() {
		for val := range numsCh {
			sqrsCh <- val * val
		}
		close(sqrsCh)
	}()

	for sqr := range sqrsCh {
		fmt.Println(sqr)
	}
}
