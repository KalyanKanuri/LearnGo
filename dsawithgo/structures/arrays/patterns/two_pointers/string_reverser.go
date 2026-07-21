package twopointers

import (
	"fmt"
)

func ReverseString(str string) {
	fmt.Println("Reverse a String")

	runes := []rune(str)
	left := 0
	right := len(runes) - 1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	fmt.Println(string(runes))
}
