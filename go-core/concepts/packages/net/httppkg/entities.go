package main

type Address struct {
	City    string
	Country string
}

type Employee struct {
	ID		int		`json:"id"`
	Name    string	`json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
	Skills  []string `json:"skills"`
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
