package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CheckAPIHealth() {
	fmt.Println("Checking Health Status of http://localhost:8080/health")
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		fmt.Println("Error getting API response: ", err)
		return
	}
	defer resp.Body.Close()

	var res any
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println("Error reading API response", err)
		return
	}

	resBytes, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		fmt.Println("Error coverting API response to JSON body", err)
		return
	}

	fmt.Printf("API Health Status: %+s\n", string(resBytes))
}
