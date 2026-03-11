package variables

import "fmt"

func VarsInGo() {
	// variable declaration and initialization in Go
	var bike string      // declaration
	bike = "Hero Xtreme" // initialization

	var speed int = 60 // declare and initialize

	max_speed := 120 // short variable declaration, only works inside functions

	fmt.Println("\n******* Variables In Go ********")
	fmt.Printf("Bike Model: %s \ncurrent speed: %d \nmax speed: %d", bike, speed, max_speed)
}
