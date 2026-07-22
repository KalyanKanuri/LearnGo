package concurrency

import "fmt"

func genNums(ch chan int) {
	ch <- 42
}

func genLetters(ch chan string) {
	ch <- "Hello"
}

func ExecuteSelect() {
	numbers := make(chan int)
	letters := make(chan string)

	go genNums(numbers)
	go genLetters(letters)

	select {
	case n := <-numbers:
		fmt.Println(n)
	case l := <-letters:
		fmt.Println(l)
	}
}
