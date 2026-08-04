package jsonpkg

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Pwd  string `json:"-"`
}

func DoMarshal() {
	emp := Employee{
		Name: "kalyan",
		Age:  26,
		Pwd:  "go.learn-26",
	}
	empJson, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
		return
	}
	fmt.Println("json.Marshal()", string(empJson))
}

func DoUnMarshal() {
	var emp Employee
	data := `
	{
		"name": "kalyan",
		"age":	26
	}
	`
	err := json.Unmarshal([]byte(data), &emp)
	if err != nil {
		fmt.Println("Error unmarshalling data", err)
		return
	}
	emp.Pwd = "go.learn-26"
	fmt.Printf("json.Unmarshal() %+v\n", emp)
}

func DoMarhsalIndent() {
	emp := Employee{
		Name: "kalyan",
		Age:  26,
		Pwd:  "go.learn-26",
	}
	empJson, err := json.MarshalIndent(emp, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json.MarshalIndent()", string(empJson))
}
