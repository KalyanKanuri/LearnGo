package datastructures

import (
	"fmt"
)

/*
Slices are nothing but dynamic arrays in Go
the only difference between an array and a slice is that
we mention capacity when creating an array whereas we don't for a slice
*/
func SlicesInGo() {
	// [3]string{"a", "b", "c"} -- this is an array
	slc := []string{"a", "b", "c", "d"}

	fmt.Println("-- Slices In Go --")
	fmt.Println(slc)

	fmt.Println("slice using make")
	/*
	make takes 3 arguments 1st -- dstype with type eg []int, 2nd -- size of ds, 3rd -- capacity of ds
	this doesn't mean the slice is limited to capacity of 10, but when the size goes beyond capacity
	go copies over all the items into a new memory object of slice with doubled capacity, we should be
	mindful of this because when we have so many records like 1billion then the memory consumption increases
	and it takes more time and memory to allocate the space and item.

	> make creates a non-empty or non-nil type while the literal ds creation gives nill in the ds
	*/
	slc2 := make([]int, 5, 10)
	fmt.Printf("int Slice %+v, size %d, capacity %d\n", slc2, len(slc2), cap(slc2))
}
