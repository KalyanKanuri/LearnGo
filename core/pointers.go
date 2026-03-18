package core

import (
	"fmt"
)

/*
 * Pointers in Go
 * --------------
 * Go is call by value only, it doesn't use references to the objects
 * so it always creates a copy of the passed value to the function, hence
 * we use pointers to update or change the actual object instead of copy
 * of the object passed, we have some exception with some types which are:
 * > slices
 * > maps
 * > interfaces
 * > funcs
 * > channels
 * -- we got a catch here, if we pass the value and update the value in func
 *    and pass it from the func it will achieve the same functionality what is done
 *    when pointers are passed but the original will not be changes, it would create
 *	  a new copy and that new copy will be returned instaed
 */

func PointersInGo() {
	fmt.Println("\n******* Pointers In Go *******")
	// declaring a pointer variable
	var ptr *int // this will be a nil pointer
	fmt.Println("zero value of ptr when declared:", ptr)

	ptr = new(int) // this will allocate memory for an int and return its address
	fmt.Println("address allocated in ptr using new: ", ptr)

	// dereferencing a pointer
	var value int = 10
	ptr = &value // using & we can access the address of a variable
	fmt.Printf("address of value: %d, is %p\n", value, ptr)
	fmt.Println("value of ptr:", *ptr) // using * we can dereference the pointer to get the value

	// modifying value through pointer
	*ptr = 20
	fmt.Printf("value of ptr after modification: %d in memory address %p\n", value, ptr) // value will be updated to 20
	// without creating a new memory, we have modified the value at the same memory address using pointer
}
