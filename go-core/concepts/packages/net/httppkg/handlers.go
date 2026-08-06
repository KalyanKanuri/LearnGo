package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var mu sync.Mutex

func rootHandler(w http.ResponseWriter, r *http.Request) {
	resp := "Hello, Backend Engineering!"
	_, err := w.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response to network", err)
		return
	}
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		emps, err := LoadEmployees("employees.json")
		if err != nil {
			fmt.Println("Error loading employees", err)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		empResp, err := json.MarshalIndent(emps, "", "	")
		if err != nil {
			fmt.Println("Error marshalling employees", err)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}
		w.Write(empResp)
	case http.MethodPost:
		mu.Lock()
		defer mu.Unlock()

		bodyBytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		fmt.Printf("Post Employees -> Request Body: %s\n", string(bodyBytes))
		if err != nil {
			fmt.Println("Error reading request body", err)
			http.Error(
				w,
				"Bad request",
				http.StatusBadRequest,
			)
			return
		}

		emps, err := LoadEmployees("employees.json")
		if err != nil {
			fmt.Println("Error loading employees", err)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		var newEmp Employee
		err = json.Unmarshal(bodyBytes, &newEmp)
		if err != nil {
			fmt.Println("Error unmarshalling new employee", err)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		emps = append(emps, newEmp)
		err = SaveEmployees("employees.json", emps)
		if err != nil {
			fmt.Println("Error Saving employees", err)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		fmt.Fprintf(w, "New Employee Created Successfully %s\n", newEmp.Name)
	}
}
