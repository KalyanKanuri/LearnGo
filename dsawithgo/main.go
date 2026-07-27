package main

import (
	twopointers "dsawithgo/structures/arrays/patterns/two_pointers"
	sdwInt "dsawithgo/structures/arrays/patterns/sliding_window"
	sdwStr "dsawithgo/structures/strings/patterns/sliding_window"
	"fmt"
)

func main() {
	fmt.Println("***** Array Patterns in DSA *****")

	fmt.Println("-- Two Pointers Pattern --")

	slc := []int{1, 2, 3, 4, 5}
	twopointers.ReverseArray(&slc)
	fmt.Println()

	palSlc := []int{4, 3, 3, 4}
	twopointers.PalindromeIdentifier(&palSlc)
	fmt.Println()

	dupSlc := []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	twopointers.RemoveDuplicates(&dupSlc)
	fmt.Println()

	str := "golang"
	twopointers.ReverseString(str)
	fmt.Println()

	elSlc := []int{3, 2, 2, 3}
	updSlc := twopointers.RemoveElement(elSlc, 3)
	fmt.Println(updSlc)

	fmt.Println("\nMove Zeroes to the end in an array")
	zeroSlc := []int{0, 1, 0, 3, 12}
	movedSlc := twopointers.MoveZeroes(zeroSlc)
	fmt.Println(movedSlc)

	fmt.Println("\nMerge two sorted arrays")
	srtArr1 := []int{1, 3, 5}
	srtArr2 := []int{2, 4, 6}
	mergedArr := twopointers.MergeSortedArrays(srtArr1, srtArr2)
	fmt.Println(mergedArr)

	fmt.Println("\nSquare And Sort an array")
	sqrSlc := []int{-4, -1, 0, 3, 10}
	sortedSqrSlc := twopointers.SortedSquares(sqrSlc)
	fmt.Println(sortedSqrSlc)

	fmt.Println("\nMax Average of a Sub Array")
	avgSlc := []int{1, 12, -5, -6, 50, 3}
	maxAvg := sdwInt.MaxAvgSubArray(avgSlc, 4)
	fmt.Printf("Max Average of a sub array in slice %+vis %.2f\n\n", avgSlc, maxAvg)

	fmt.Println("***** String Patterns in Go *****")
	subStr := "abcabca"
	maxLen := sdwStr.LongestSubString1(subStr)
	fmt.Printf("Max substring length of string %s is %d\n", subStr, maxLen)
	maxLen = sdwStr.LongestSubStr2(subStr)
	fmt.Println("Max Sub Str with last idx -- approach 2")
	fmt.Println(maxLen)
}
