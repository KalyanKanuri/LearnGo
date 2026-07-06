package funcs

import (
	"fmt"
)

// Closures are anyonymous functions that can be assigned to a variable and can be called later. Closures can access and modify variables defined outside of their scope.
func Closures() {
	square := func(n int) {
		fmt.Printf("The square of %d is %d\n", n, n*n)
	}
	square(5)
}