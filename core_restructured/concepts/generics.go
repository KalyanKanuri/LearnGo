package concepts

import "fmt"

/* Generics:
   --------
   > generics allows us to write structs, funcs, interfaces to work with multiple types without risking type safety
   > without generics we have to write different funcs for each type
*/

type float interface {
	float32 | float64
}

// --------- Without Generics ---------------- //
func MaxInt(a, b int) {
	if a > b {
		fmt.Println(a, "is max")
	} else {
		fmt.Println(b, "is max")
	}
}

func MaxFloat(a, b float64) {
	if a > b {
		fmt.Println(a, "is max")
	} else {
		fmt.Println(b, "is max")
	}
}

// ---------------- With Generics ----------------- //
func calculateMax[T int | float64](a, b T) {
	if a > b {
		fmt.Println(a, "is max")
	} else {
		fmt.Println(b, "is max")
	}
}

func GenericsInGo() {
	calculateMax(5, 6)
	calculateMax(5.5, 6.5)
}
