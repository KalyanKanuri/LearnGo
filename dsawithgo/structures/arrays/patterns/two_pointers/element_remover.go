package twopointers

import (
	"fmt"
)

func RemoveElement[T comparable](arr []T, val T) []T {
	fmt.Println("Remove given element from array")

	writer := 0
	for reader := range arr {
		if arr[reader] != val {
			arr[writer] = arr[reader]
			writer++
		}
	}
	return arr[:writer]
}
