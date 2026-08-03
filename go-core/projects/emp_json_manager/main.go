package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Employee struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address Address  `json:"address"`
	Skills  []string `json:"skills"`
}

func main() {
	jsonData, err := os.ReadFile("employees.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var emps []Employee
	err = json.Unmarshal(jsonData, &emps)
	if err != nil {
		fmt.Println("Error unmarshaling json data", err)
		return
	}

	fmt.Println("============ Employee Summary ============")
	for i, emp := range emps {
		fmt.Printf("%d. %s\nAge: %d\nCity: %s\nCountry: %s\nSkills: %s\n\n",
			i+1, emp.Name, emp.Age, emp.Address.City, emp.Address.Country, emp.Skills[:],
		)
	}
	fmt.Println("==========================================")

	newEmp := Employee{
		Name: "Das",
		Age:  29,
		Address: Address{
			City:    "Hyderabad",
			Country: "India",
		},
		Skills: []string{"React", "Next JS"},
	}
	emps = append(emps, newEmp)

	empsJson, err := json.MarshalIndent(emps, "", "	")
	if err != nil {
		fmt.Println("Error marshalling json", err)
		return
	}
	err = os.WriteFile("employees.json", empsJson, 0777)
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}
}
