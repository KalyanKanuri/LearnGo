package main

import (
	"encoding/json"
	"os"
)

func LoadEmployees(filename string) ([]Employee, error) {
	var emps []Employee
	empBytes, err := os.ReadFile(filename)
	if err != nil {
		return []Employee{}, err
	}

	err = json.Unmarshal(empBytes, &emps)
	if err != nil {
		return []Employee{}, err
	}
	return emps, err
}

func SaveEmployees(filename string, emps []Employee) error {
	empBytes, err := json.MarshalIndent(emps, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, empBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
