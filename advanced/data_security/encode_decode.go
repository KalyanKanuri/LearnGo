package datasecurity

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Toy struct {
	ID      int
	Name    string
	Bought  int
	InStock int
}

var payload = `{"ID":1,"Name":"Hot Wheels","Bought":25,"InStock":35}`

/*
 * Encode Decode in Go
 * -------------------
 * The only difference between Marshal_Unmarshal & Encode_Decode in Go is that
 * Marshal_Unmarshal returns the data whereas Encode_Decode takes in a writer
 * or reader instead of returning the data so that we can have the felxibility
 * to write the data anywhere or read the data from anywhere.
 */
func EncodeDecodeInGo() {
	fmt.Println("-- Encoding & Decoding data in Go --")
	fmt.Print("Encoder Result: ")
	encoder := json.NewEncoder(os.Stdout)

	toy := Toy{
		ID:      1,
		Name:    "Hot Wheels",
		Bought:  25,
		InStock: 35,
	}
	encoder.Encode(toy)

	decoder := json.NewDecoder(strings.NewReader(payload))

	var toy2 Toy
	decoder.Decode(&toy2)

	fmt.Printf("Decoder Result: %+v\n", toy2)
}
