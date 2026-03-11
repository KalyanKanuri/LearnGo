package main

import "fmt"
import "learning/go/basics"
import "learning/go/projects/glogger"

func main() {
	// Printing in Go
	print("Started Learning Go")
	print("default print func from go runtime doesn't add new line")
	println("default println func from go runtime adds new line")
	fmt.Print("print func from fmt, provides additional formatting")
	fmt.Println("println func from fmt, provides new line with formatting")
	fmt.Printf("printf func from fmt, provides control over vars while printing\n")

	// Variables in Go
	basics.VarsInGo()

	// Constants in Go
	basics.ConstInGo()
	basics.IotaInGO()

	// 1st project custom logger
	fmt.Println("******* Custom Logger *******")
	glog.Log(0, "implemented custom logger")
	glog.Log(1, "used iota for log levels")
	glog.Log(2, "defined a type Loglevel")
	glog.Log(3, "created a func Log")
	glog.Log(4, "accepts LogLevel and a log msg")
}
