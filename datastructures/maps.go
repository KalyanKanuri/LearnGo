package datastructures

import (
	"fmt"
)

/*
basic definition: map[keyType]valueType
*/
func MapsInGo() {
	fmt.Println("\n******* Maps In Go *******")
	// using var
	// this will result in nill[zero-value]
	//```go
	// var student map[string]string
	// ```
	// the above method is not suggested as it will just decalres and doesn't allocate any memory

	// using literal
	// insert operation
	fmt.Println("-- using literal method --")
	studentDetails := map[string]string{
		"Name":            "Pavitr Prabhakar",
		"Age":             "20",
		"Address":         "Mumbai, India",
		"Class":           "2nd Year",
		"Speciailization": "Computer Science",
		"University":      "MCU",
	}

	// access and read
	key := "Name"
	// comma-ok idiom
	if value, ok := studentDetails[key]; ok {
		fmt.Printf("Key %s, Value %s, Ok %t\n", key, value, ok)
	}

	// using make -- 2 level map
	// insert
	fmt.Println("-- using make func --")
	appConfig := make(map[string]map[string]string) // this will create only one mape
	appConfig["App"] = make(map[string]string)      // we have to make again the inner map
	appConfig["App"]["AppName"] = "QuickByte"
	appConfig["App"]["version"] = "1.0.0"
	appConfig["App"]["backend"] = "Go"
	appConfig["App"]["frontend"] = "React"
	appConfig["App"]["some giberis"] = "giberish"

	fmt.Printf("%+v\n", appConfig["App"])

	delete(appConfig["App"], "some giberis")

	fmt.Println("after delete operation")
	fmt.Printf("%+v\n", appConfig["App"])

	fmt.Println("-- range loop in map --")
	for key, value := range appConfig["App"] {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
