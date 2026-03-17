package core

import "fmt"

func simulatePanic(panicFlag bool) {
	if panicFlag {
		panic("Some Error occured")
	}
}

func PanicRecoverInGo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}
	}()
	simulatePanic(true)
}
