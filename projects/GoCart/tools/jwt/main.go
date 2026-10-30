package main

import (
	"flag"
	"fmt"
)

func main() {
	jwtAlgo := flag.String("algo", "HS256", "JWT algorithm: HS256 (32 bytes), HS384 (48 bytes) or HS512 (64 bytes)")
	flag.Parse()

	secretHex, err := GenSecret(*jwtAlgo)
	if err != nil {
		fmt.Println("error generating secret:", err)
		return
	}

	fmt.Println("Your secure secret key:", secretHex)
}
