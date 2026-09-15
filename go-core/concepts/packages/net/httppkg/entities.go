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

type EmployeeRequest struct {
	Name    string   `json:"name" validate:"required"`
	Age     *int     `json:"age"`
	Address Address  `json:"address"`
	Skills  []string `json:"skills"`
}

type ValidationErrorResponse struct {
	Errors []FieldError `json:"errors"`
}
