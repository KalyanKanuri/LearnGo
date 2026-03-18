package core

import (
	"errors"
	"fmt"
)

// Factorial using recurssion
func recursionInGo(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursionInGo(n-1)
}

// Closures -- returning a func from a func
func intSeq() func() int {
	i := 0
	// this is also called as an anonymous function
	return func() int {
		i++
		return i
	}
}

// Variadic Function
func sum(nums ...int) int {
	total := 0
	if len(nums) > 0 {
		for _, num := range nums {
			total += num
		}
	}
	return total
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot perform division")
	}
	return a / b, nil
}

func simulateDefer() {
	fmt.Println("simulateDefer: start")
	defer fmt.Println("simulateDefer: middle")
	//eventhough we are expecting this print at end
	// defer will make the above statment to print last
	// if we have two defers it will follow lifo principle that is the last defer will execute first
	fmt.Println("simulateDefer: end")
	fmt.Println("multi defer in func")
	defer fmt.Println("simulateDefer: second defer")
}

func FuncsInGo() {
	// Functions in Go
	fmt.Println("\n******* Functions In Go *******")
	/* Functions in GO
	 * ---------------
	 * from the start onwards we are working with functions while learning GO
	 * but we can do much more with GO
	 * > calling the function within itself -- recursion
	 * > returning a function from a function -- closure
	 * > assigning a function to a variable -- this var acts like an alias to the function
	 * > ability to take custom or optional arguements -- variadic function
	 * > return multi values
	 */

	fmt.Println("-- Recursion --")
	val := recursionInGo(5)
	fmt.Println("recursion value:", val)

	fmt.Println("-- Closures --")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("-- Function as variable --")
	logger := func(msg string) {
		fmt.Println(msg)
	}
	logger("using logger to work as func while declared as var")

	fmt.Println("-- Variadic Function --")
	// if we don't pass any nums the total will simply be 0
	sumOfNums := sum(1, 2, 3, 4)
	fmt.Println("sum of numbers:", sumOfNums)

	fmt.Println("-- Return multi values --")
	if val, err := divide(4, 2); err == nil {
		fmt.Println("division result is:", val)
	} else {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("-- defer in Go --")
	simulateDefer()

	fmt.Println("-- Anonymous Functions in Go --")
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println("AF: calling sum ->", sum(5, 2))
}
