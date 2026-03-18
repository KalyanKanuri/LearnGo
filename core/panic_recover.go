package core

import "fmt"

func simulatePanic(panicFlag bool) {
	if panicFlag {
		panic("--> Simulate panic in go")
	}
}

func PanicRecoverInGo() {
	fmt.Println("-- Panic & Recover in Go --")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}
	}()
	simulatePanic(true)
}
