package concurrency

import "fmt"

func ExecuteDefault() {
	ch := make(chan int, 3)

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("No value available")
	}
}
