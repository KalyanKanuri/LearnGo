package concurrency

import "fmt"

func GenerateNums() {
	ch := make(chan int)
	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("Producer finished. Consumer exiting.")
}
