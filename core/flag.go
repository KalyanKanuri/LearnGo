package core

import (
	"flag"
	"fmt"
)

func FlagPkgInGo() {
	fmt.Println("-- Flag Pkg in Go --")
	name := flag.String("name", "", "name to greet")
	age := flag.Int("age", 0, "age to provide")

	flag.Parse()

	fmt.Printf("Hello %s, Age: %d", *name, *age)
}
