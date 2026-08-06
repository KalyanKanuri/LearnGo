package main

type Address struct {
	City    string
	Country string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
	Skills  []string
}
