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
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	if err != nil {
		fmt.Println("Error writing response to network", err)
		return
	}
	// this is to understand how recovery middleware works uncomment below to visualize
	//panic("Intentional panic for testing RecoveryMiddleware")
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
		w.WriteHeader(200)
		w.Write(empResp)
	case http.MethodPost:
		bodyBytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			fmt.Println("Error reading request body", err)
			http.Error(
				w,
				"Bad request",
				http.StatusBadRequest,
			)
			return
		}
		fmt.Printf("Post Employees -> Request Body: %s\n", string(bodyBytes))

		var empReq EmployeeRequest
		err = json.Unmarshal(bodyBytes, &empReq)
		if err != nil {
			fmt.Println("Error unmarshalling new employee", err)
			http.Error(
				w,
				"Bad Request",
				http.StatusBadRequest,
			)
			return
		}

		err = ValidateEmployeeRequest(empReq)
		if err != nil {
			fmt.Println("Error in Request body", err)
			http.Error(
				w,
				"Bad Request",
				http.StatusBadRequest,
			)
			return
		}

		mu.Lock()
		defer mu.Unlock()
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

		newEmp := toEmployee(empReq)
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

		w.WriteHeader(201)
		fmt.Fprintf(w, "New Employee Created Successfully %s\n", newEmp.Name)
	}
}
