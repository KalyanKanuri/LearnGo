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
					w.Header().Set(
						"content-type", "application/json",
					)
					w.WriteHeader(http.StatusInternalServerError)
					_, err = w.Write([]byte(`
						{
							"error": "Internal Server Error"
						}`),
					)
					if err != nil {
						fmt.Println("Error writing response to network", err)
					}
					return
				}

				w.Header().Set(
					"content-type", "application/json",
				)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errResp))
				return
			}
			fmt.Println("Error in Request body", err)
			w.Header().Set(
				"content-type", "application/json",
			)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`
				{
					"error": "Bad Request"
				}`),
			)
			return
		}

		newEmp := toEmployee(empReq)

		mu.Lock()
		defer mu.Unlock()
		emps, err := LoadEmployees("employees.json")
		if err != nil {
			fmt.Println("Error loading employees", err)
			w.Header().Set(
				"content-type", "application/json",
			)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`
				{
					"error": "Internal Server Error"
				}`),
			)
			if err != nil {
				fmt.Println("Error writing response to network", err)
			}
			return
		}

		newEmp.ID = nextEmployeeID(emps)

		emps = append(emps, newEmp)
		err = SaveEmployees("employees.json", emps)
		if err != nil {
			fmt.Println("Error Saving employees", err)
			w.Header().Set(
				"content-type", "application/json",
			)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`
				{
					"error": "Internal Server Error"
				}`),
			)
			if err != nil {
				fmt.Println("Error writing response to network", err)
			}
			return
		}

		
		resp, err := json.MarshalIndent(newEmp, "", " ")
		if err != nil {
			fmt.Println("Error marshalling response", err)
			w.Header().Set(
				"content-type", "application/json",
			)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write([]byte(`
				{
					"error": "Internal Server Error"
				}`),
			)
			if err != nil {
				fmt.Println("Error writing response to network", err)
			}
			return
		}
		w.Header().Set(
			"content-type", "application/json",
		)
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(resp)
		if err != nil {
			fmt.Println("Error writing response to network", err)
		}
	}
}
