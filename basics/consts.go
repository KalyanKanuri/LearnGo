package basics

import "fmt"

/* Constants in Go
 * ---------------
 * Constants cannot be changed once assigned a value
 * these should be literals, cannot assign at runtime
 * these are decided as constants at runtime
 */

func ConstInGo() {
	const PI = 3.14
	const HOST = "localhost"
	const PORT = "8080"
	const MAX_CONNECTIONS = 4

	fmt.Println("\n******* Constants In Go *******")
	fmt.Println(HOST, PORT, MAX_CONNECTIONS)
}
