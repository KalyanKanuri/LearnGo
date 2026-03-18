package core

import (
	"fmt"
)

type Number interface {
	int | float64 | float32
}

func doSum[T Number](nums ...T) T {
	var total T
	for _, val := range nums {
		total += val
	}
	return total
}

/*
 * Generics in Go
 * --------------
 * we have generics in every programming language to support multiple types
 * as go is a typed language it enforces us to use a type unlike python or js
 * so instead of capable to deal with only a single type to pass args or return
 * from a func we can use types to expect any kind of type within defined interface
 * e.g., In below example we can see that  we have defined a type combined with int,
 * float this enforces the passed or returned arguement to be within these any of these two
 */
func GenericsInGo() {
	fmt.Println("\n******* Generics in Go *******")
	// as we have defined sum to accept any type within int and float
	// we are able to pass int as well float as args
	total := doSum(1, 2.5, 10, 30.6, 12)
	fmt.Printf("total: %+v", total)
}
