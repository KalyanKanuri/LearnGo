package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	resp := "Hello, Backend Engineering!"
	bytesWritten, err := w.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response to network", err)
		return
	}
	fmt.Printf("%d, bytes written into response\n", bytesWritten)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server start up failed", err)
		return
	}
}
