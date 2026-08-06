package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/employees", employeeHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server start up failed", err)
		return
	}
}
