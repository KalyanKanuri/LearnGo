package main

import (
	"encoding/json"
	"os"
)

func LoadEmployees(filename string) ([]byte, error) {
	empBytes, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return empBytes, err
}

func SaveEmployees(filename string, emps []Employee) ([]byte, error) {
	empBytes, err := json.MarshalIndent(emps, "", "	")
	if err != nil {
		return empBytes, err
	}

	err = os.WriteFile(filename, empBytes, 0644)
	if err != nil {
		return []byte{}, err
	}
	return empBytes, nil
}
