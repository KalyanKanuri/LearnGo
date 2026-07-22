package concurrency

import "fmt"

func GenerateSquares() {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			n := i + 1
			ch <- n * n
		}
		close(ch)
	}()

	for sqr := range ch {
		fmt.Println(sqr)
	}
}
