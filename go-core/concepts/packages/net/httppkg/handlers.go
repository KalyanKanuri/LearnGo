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
	resp := "Hello, Backend Engineer!"
	_, err := w.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response to network", err)
		http.Error(
			w,
			"Internal Server Error",
			http.StatusInternalServerError,
		)
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

		empResp, err := json.MarshalIndent(emps, "", " ")
		if err != nil {
			fmt.Println("Error marshalling employees", err)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(empResp)
	case http.MethodPost:
		reqDecoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var empReq EmployeeRequest
		err := reqDecoder.Decode(&empReq)
		if err != nil {
			fmt.Println("Error Decoding request body", err)
			http.Error(
				w,
				"Bad Request",
				http.StatusBadRequest,
			)
			return
		}

		var extra any
		err = reqDecoder.Decode(&extra)
		if err != io.EOF {
			fmt.Println("Invalid Request Body", extra)
			http.Error(
				w,
				"Bad Request",
				http.StatusBadRequest,
			)
			return
		}

		err = ValidateEmployeeRequest(empReq)
		if err != nil {
			if validationErrs, ok := err.(ValidationErrors); ok {
				errs := ValidationErrorResponse{
					Errors: validationErrs.Errs,
				}

				errResp, err := json.MarshalIndent(errs, "", " ")
				if err != nil {
					http.Error(
						w,
						"Internal Server Error",
						http.StatusInternalServerError,
					)
					return
				}

				http.Error(
					w,
					string(errResp),
					http.StatusBadRequest,
				)
				return
			}
			fmt.Println("Error in Request body", err)
			http.Error(
				w,
				"Bad Request",
				http.StatusBadRequest,
			)
			return
		}

		newEmp := toEmployee(empReq)

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

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "New Employee Created Successfully %s\n", newEmp.Name)
	}
}
