package main

import (
	"fmt"
	twopointers "revisedsa/two_pointers"
)

func main() {
	fmt.Println("Reverse an Array")
	revArr := []int{1,2,3,4,5}
	revdArr := twopointers.ReverseArray(revArr)
	fmt.Println(revdArr)

	fmt.Println("Remove Duplicates from a Sorted Array")
	dupsArr := []int{1,1,2,2,3,3,4,5,6,7,7}
	dupdArr := twopointers.RemoveDuplicates(dupsArr)
	fmt.Println(dupdArr)

	fmt.Println("Max Average Sub Array")
	twoSum := []int{1, 2, 3, 4, 5}
	target := 5
	pointers := twopointers.TwoSum(twoSum, target)
	fmt.Println(pointers)
}
