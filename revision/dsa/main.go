package main

import (
	"fmt"
	twopointers "revisedsa/two_pointers"
	threepointers "revisedsa/three_pointers"
)

func main() {
	fmt.Println("Reverse an Array")
	revArr := []int{1, 2, 3, 4, 5}
	revdArr := twopointers.ReverseArray(revArr)
	fmt.Println(revdArr)

	fmt.Println("Remove Duplicates from a Sorted Array")
	dupsArr := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 7, 7}
	dupdArr := twopointers.RemoveDuplicates(dupsArr)
	fmt.Println(dupdArr)

	fmt.Println("Max Average Sub Array")
	twoSum := []int{1, 2, 3, 4, 5}
	target := 5
	pointers := twopointers.TwoSum(twoSum, target)
	fmt.Println(pointers)

	fmt.Println("Merge Two Sorted Arrays")
	mrg1 := []int{1, 3, 5, 7}
	mrg2 := []int{2, 4, 6, 8}
	mrgd := twopointers.MergeSortedArrays(mrg1, mrg2)
	fmt.Println(mrgd)

	fmt.Println("Move Zeroes in an Array")
	zeroes := []int{0, 1, 0, 2, 0, 3, 4, 5, 6, 0}
	mvdZrs := twopointers.MoveZeroes(zeroes)
	fmt.Println(mvdZrs)

	fmt.Println("Merge 2 Sorted Arrays with O(1)")
	mrg_1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	mrg_2 := []int{2, 5, 6}
	n := 3
	threepointers.Merge2SortedArrs(mrg_1, mrg_2, m, n)

	fmt.Println("Square and Sort and array")
	sqrs := []int{-4,-2,0,5,6}
	sqrd := threepointers.SortSquares(sqrs)
	fmt.Println(sqrd)
}
