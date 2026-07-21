package twopointers

import (
	"fmt"
)

func RemoveDuplicates[T comparable](arr *[]T) {
	fmt.Println("Remove duplicates from an array")

	writer := 1
	for reader := range *arr {
		if reader == 0 {
			continue
		}
		if (*arr)[reader] != (*arr)[reader -1] {
			(*arr)[writer] = (*arr)[reader]
			writer++
		}
	}
	uniqueSlc := (*arr)[:writer]
	fmt.Println(uniqueSlc)
}