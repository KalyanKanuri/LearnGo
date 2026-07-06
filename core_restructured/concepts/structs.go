package concepts

import (
	"fmt"
)

/* Structs are used to group together data of different types into a single entity.
> structs are similar to classes in other programming languages, but they do not have methods or inheritance.
> structs are defined using the `type` keyword, followed by the name of the struct and the `struct` keyword.

e.g., type Person struct {
	Name string
	Age  int
}
*/

// defining a struct
type StructExample struct {
	Name      string
	NumFields int
	isValid   bool
}

func StructsInGo() {
	// initializing a struct
	se := StructExample{
		Name:      "Struct Example",
		NumFields: 3,
		isValid:   true,
	}

	fmt.Println("Struct Example:", se)

	// Accessing struct fields
	fmt.Println("Name:", se.Name)
	fmt.Println("NumFields:", se.NumFields)
	fmt.Println("isValid:", se.isValid)

	se.UpdateName("updatedStructName")
	fmt.Println("Updated Name:", se.Name)

	fmt.Println("--- composite structs ---")
	CompositionExample()
}

// these type of funcs are called as methods in Go. similar to methods in classes in other programming languages.
// these methods can be accessed with the struct instance. e.g., se.UpdateName("New Name")
func (se *StructExample) UpdateName(newName string) {
	se.Name = newName
}

// composition via struct emebedding
// composition means building an object embedded with other objects.
// see example below where we build a car object with Engine, wheels and Model as embedded objects.
type Engine struct {
	Type       string
	Horsepower int
}

func (e Engine) Start() {
	fmt.Println("Engine started")
}

type Wheels struct {
	Type  string
	Count int
}

func (w Wheels) Rotate() {
	fmt.Println("Wheels rotating")
}

type Car struct {
	Model string
	Engine
	Wheels
}

func CompositionExample() {
	car := Car{
		Model: "Toyota Camry",
		Engine: Engine{
			Type:       "V6",
			Horsepower: 301,
		},
		Wheels: Wheels{
			Count: 4,
			Type:  "Alloy",
		},
	}
	fmt.Printf("Car: %+v\n", car)
	// Method Promotion: we can call the methods of the embedded struct directly on the outer struct.
	car.Start()  // this calls the Start method of the Engine struct
	car.Rotate() // this calls the Rotate method of the Wheels struct
}
