package funcs

import (
	"fmt"
)

// variadic function is a function that can take zero or more arguments of a specified type.
func VariadicFunc(args ...int) {
	fmt.Println("Number of arguments:", len(args))

	for i, arg := range args {
		fmt.Printf("Argument %d: %d\n", i+1, arg)
	}
}
