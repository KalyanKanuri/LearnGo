package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	resp := "Hello, Backend Engineering!"
	_, err := w.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response to network", err)
		return
	}
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	var emps []Employee
	empBytes, err := LoadEmployees("employees.json")
	if err != nil {
		fmt.Println("Error loading employees", err)
		w.Write([]byte("Error loading employees"))
		return
	}

	err = json.Unmarshal(empBytes, &emps)
	if err != nil {
		fmt.Println("Error loading employees", err)
		w.Write([]byte("Error loading employees"))
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.Write(empBytes)
	case http.MethodPost:
		bodyBytes, err := io.ReadAll(r.Body)
		fmt.Printf("Post Employees -> Request Body: %s\n", string(bodyBytes))
		if err != nil {
			fmt.Println("Error reading request body", err)
			w.Write([]byte("Error reading request body"))
			return
		}
		var newEmp Employee
		err = json.Unmarshal(bodyBytes, &newEmp)
		if err != nil {
			fmt.Println("Error unmarshalling json", err)
			w.Write([]byte("Error unmarshalling json"))
			return
		}
		emps = append(emps, newEmp)
		newEmpBytes, err := SaveEmployees("employees.json", emps)
		if err != nil {
			fmt.Println("Error Saving Employees", err)
			w.Write([]byte("Error saving employees"))
			return
		}
		w.Write(newEmpBytes)
	}
}
