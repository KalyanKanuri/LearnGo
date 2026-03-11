package main;

import "fmt"
import "learning/go/variables"

func main() {
	// Printing in Go
	print("Started Learning Go")
	print("default print func from go runtime doesn't add new line")
	println("default println func from go runtime adds new line")
	fmt.Print("print func from fmt, provides additional formatting")
	fmt.Println("println func from fmt, provides new line with formatting")
	fmt.Printf("printf func from fmt, provides control over vars while printing\n")

	// Variables in Go
	variables.VarsInGo()
}