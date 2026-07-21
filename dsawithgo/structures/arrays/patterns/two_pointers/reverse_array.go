package twopointers

import (
	"fmt"
)

func ReverseArray[T any](arr *[]T) {
	fmt.Println("Reversing an array using 2 pointers")
	left := 0
	right := len(*arr) - 1

	for left < right {
		(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
		left++
		right--
	}
	fmt.Println(*arr)
}
