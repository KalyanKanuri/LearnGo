package basics

import "fmt"

const (
	SUNDAY = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

/*
automatically completes the continuation when used iota
Go doesn't have any keyword like ENUM but we can use iota to implement the features of ENUM
*/
func IotaInGO() {
	fmt.Println("-- IOTA in Go --")
	fmt.Println(SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY)
}
