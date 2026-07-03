package concepts

import (
	"fmt"
)

func SlicesInGo() {
	// slice declaration
	var slc []int
	// This is a nil slice that means the slice object is not created in memory as there's no fixed size

	// slice value infusion
	// To infuse items into a slice we have to either create the slice with make() func or grow the slice with append() func
	// RULE OF THUMB - when we want to replace an item in the slice at an existing index then use make() if we want to push a new value into the slice use append()
	slc = make([]int, 5) // we are creating a slice of size 5, this is note fixed unlike arrays
	slc[1] = 5

	slc = append(slc, 10) // this will grow the slice from size 5 to 6 and append value 10 at last position

	fmt.Println(slc)
}