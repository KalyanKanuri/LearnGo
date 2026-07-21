package twopointers

import (
	"fmt"
)

func PalindromeIdentifier[T comparable](arr *[]T) {
	fmt.Println("Palindrome identification using two pointers")
	left := 0
	right := len(*arr) - 1

	for left < right {
		if (*arr)[left] == (*arr)[right] {
			left++
			right--
		} else {
			// mismatch found not a palindrome
			fmt.Printf("Mismatch found returning -- arr[%d] != arr[%d] vals %+v != %+v\n", left, right, (*arr)[left], (*arr)[right])
			return
		}
	}
	fmt.Printf("Palindrome identified %+v\n", *arr)
}
