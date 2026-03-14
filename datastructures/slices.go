package datastructures

import (
	"fmt"
	"slices"
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

	// append operation
	slc2 = append(slc2, 1, 2, 3, 4, 5)
	fmt.Printf("after append operation: %+v, size %d, capacity %d\n", slc2, len(slc2), cap(slc2))

	// slices-package find index of an element in a slice
	fmt.Println("-- using slices package --")
	fmt.Println("> slices.Index()")
	idx := slices.Index(slc2, 5)
	fmt.Printf("index of 5 in slc2: %d\n", idx)

	// slices-package insert at certain position
	fmt.Println("> slices.Insert()")
	slc2 = slices.Insert(slc2, 2, 100)
	fmt.Printf("after insert operation: %+v, size %d, capacity %d\n", slc2, len(slc2), cap(slc2))

	// slices-package delete at certain position
	fmt.Println("> slices.Delete()")
	slc2 = slices.Delete(slc2, 2, 3) // this will delete the element at index 2 and 3
	fmt.Printf("after delete operation: %+v, size %d, capacity %d\n", slc2, len(slc2), cap(slc2))

	// slices-package check if two slices are equal
	fmt.Println("> slices.Equal()")
	equal := slices.Equal(slc2, []int{0, 0, 1, 2, 3, 4, 5}) // this will return true if both the slices are equal
	fmt.Printf("are the slices equal? %t\n", equal)

	// slices-package reverse a slice
	fmt.Println("> slices.Reverse()")
	slices.Reverse(slc2)
	fmt.Printf("after reverse operation: %+v, size %d, capacity %d\n", slc2, len(slc2), cap(slc2))

	// slices-package sort a slice
	fmt.Println("> slices.Sort()")
	slices.Sort(slc2)
	fmt.Printf("after sort operation: %+v, size %d, capacity %d\n", slc2, len(slc2), cap(slc2))

	// slices-package clone a slice
	fmt.Println("> slices.Clone()")
	cloned := slices.Clone(slc2)
	fmt.Printf("after clone operation: %+v, size %d, capacity %d\n", cloned, len(cloned), cap(cloned))

	// slices-package Repeat a slice
	fmt.Println("> slices.Repeat()")
	repeated := slices.Repeat([]int{1, 2, 3}, 3)
	fmt.Printf("after repeat operation: %+v, size %d, capacity %d\n", repeated, len(repeated), cap(repeated))

	// slices-package Compact a slice
	fmt.Println("> slices.Compact()")
	compact := slices.Compact([]int{1, 2, 0, 0, 3, 0, 0, 5, 0}) // help to remove consequtive repeated values
	fmt.Printf("after compact operation: %+v, size %d, capacity %d\n", compact, len(compact), cap(compact))

	// slices-package DeleteFunc a slice
	fmt.Println("> slices.DeleteFunc()")
	deleteFunc := slices.DeleteFunc([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i%2 == 0 // this will delete all the even numbers from the slice
	})
	fmt.Printf("after delete func operation: %+v, size %d, capacity %d\n", deleteFunc, len(deleteFunc), cap(deleteFunc))

	// slices-package EqualFunc a slice
	fmt.Println("> slices.EqualFunc()")
	equalFunc := slices.EqualFunc([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, func(i, j int) bool {
		return i == j
	})
	fmt.Printf("are the slices equal? %t\n", equalFunc)
}
