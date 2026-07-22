package concepts

import (
	"fmt"
)

func MapsInGo() {
	// map declaration
	var map1 map[string]string
	// this is just the declaration, this will not create the map object in memory it's a nil map

	// map init
	map1 = make(map[string]string)
	map1["key1"] = "value1"

	// comma-ok idiom
	// when a key is non existent in map, go runtime will not error out but assigns a zero-value to the key
	// this is where we use comma-ok idiom this will check if a key is present in the map or not and return a boolean

	value, ok := map1["key2"]
	fmt.Println(value) // result - white space as string zero-value is white space
	fmt.Println(ok)    // result - false as the key is not present in map
	// we have to use this comma-ok idiom in below way
	if value, ok := map1["key1"]; ok {
		// this will be executed only if the key is present.
		fmt.Println(value)
	}
	fmt.Println(map1)
}
