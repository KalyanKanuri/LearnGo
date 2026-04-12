package concurrency

import (
	"fmt"
)

func generator(done chan struct{}) <-chan int {
	nums := make(chan int)

	go func() {
		defer close(nums)
		for i := range 10 {
			fmt.Println("generated", i+1)
			select {
			case <-done:
				return
			case nums <- i + 1:
			}
		}
	}()

	return nums
}

func doubler(done <-chan struct{}, nums <-chan int) <-chan int {
	doubles := make(chan int)

	go func() {
		defer close(doubles)
		for num := range nums {
			fmt.Println("doubled", num+num)
			select {
			case doubles <- num + num:
			case <-done:
				return
			}
		}
	}()

	return doubles
}

func multiplier(done <-chan struct{}, doubles <-chan int) <-chan int {
	multiples := make(chan int)

	go func() {
		defer close(multiples)
		for num := range doubles {
			select {
			case multiples <- num * num:
				continue
			case <-done:
				return
			}
		}
	}()

	return multiples
}

func PipelinePatternGo() {
	fmt.Println("-- Pipeline pattern in Go --")
	done := make(chan struct{})

	nums := generator(done)
	doubles := doubler(done, nums)
	multiples := multiplier(done, doubles)

	for multiple := range multiples {
		fmt.Println("multiplier:", multiple)
	}
}
