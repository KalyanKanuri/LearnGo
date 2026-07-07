package funcs

import (
	"fmt"
)

// Anonymous function is a function that is defined without a name. It can be assigned to a variable and called later.
func AnonymousFunc() {
	func() {
		fmt.Println("This is an anonymous function.")
	}()
}
