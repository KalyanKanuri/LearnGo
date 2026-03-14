package datastructures

import (
	"fmt"
)

/*
Arrays are of fixed sizes in Go
*/
func ArraysInGo() {
	fmt.Println("-- Arrays in Go --")
	// when we mention size then only go will consider it as array else as a slice
	arr := [5]int{1, 2, 3, 4, 5}
	// arr = append(arr, 6)
	// we can't add an item to a fixed array
	fmt.Println(arr)
}
