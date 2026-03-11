package basics

import "fmt"

func VarsInGo() {
	// variable declaration and initialization in Go
	var bike string      // declaration
	bike = "Hero Xtreme" // initialization

	var speed int = 60 // declare and initialize

	// := this infers the type at runtime
	max_speed := 150 // short variable declaration, only works inside functions

	fmt.Println("\n******* Variables In Go ********")
	fmt.Printf("Bike Model: %s \ncurrent speed: %d \nmax speed: %d\n", bike, speed, max_speed)

	/* zero-value
	 * ----------
	 * In GO, every initialized variable will be assigned with a zero-value
	 * in this way, GO ensures that there's no unintialized value left over
	 */
	var (
		a int
		b string
		c bool
		d []int
		e map[string]int
		f *float64
	)
	fmt.Println("-- Zero values in Go --")
	fmt.Println(a, b, c, d, e, f)
	// Output: 0  false [] map[] <nil>
	// these are the default zero-values for above declared different types of vars
}
