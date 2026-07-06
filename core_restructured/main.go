package main

import (
	"coreconcepts/concepts"
	"coreconcepts/concepts/funcs"
	"fmt"
)

func main() {
	fmt.Println("***** Arrays In Go *****")
	concepts.ArraysInGo()
	fmt.Println()

	fmt.Println("***** Slices In Go *****")
	concepts.SlicesInGo()
	fmt.Println()

	fmt.Println("***** Maps In Go *****")
	concepts.MapsInGo()
	fmt.Println()

	fmt.Println("***** Structs In Go *****")
	concepts.StructsInGo()
	fmt.Println()

	fmt.Println("***** GoTo Statement in Go *****")
	concepts.GoToStatement()
	fmt.Println()

	fmt.Println("***** Variadic Functions In Go *****")
	funcs.VariadicFunc(1, 2, 3, 4, 5)
	funcs.VariadicFunc()
	fmt.Println()

	fmt.Println("***** Anonymous Functions In Go *****")
	funcs.AnonymousFunc()
	fmt.Println()

	fmt.Println("***** Closures In Go *****")
	funcs.Closures()
	fmt.Println("")
}
