package main

import "strings"

type FieldError struct {
	Field   string
	Message string
}

type ValidationErrors struct {
	Errs []FieldError
}

func (ve ValidationErrors) Error() string {
	if len(ve.Errs) > 0 {
		var sb strings.Builder

		for _, fe := range ve.Errs {
			sb.WriteString(fe.Field)
			sb.WriteString(": ")
			sb.WriteString(fe.Message)
			sb.WriteString("; ")
		}
		return sb.String()
	}

	return ""
}

func ValidateEmployeeRequest(req EmployeeRequest) error {
	var fes []FieldError

	if req.Name == "" {
		err := FieldError{
			Field:   "name",
			Message: "required",
		}
		fes = append(fes, err)
	}

	if req.Age != nil {
		if *req.Age < 18 {
			err := FieldError{
				Field:   "age",
				Message: "must be atleast 18",
			}
			fes = append(fes, err)
		}
	} else {
		err := FieldError{
			Field:   "age",
			Message: "required",
		}
		fes = append(fes, err)
	}

	if req.Address.City == "" {
		err := FieldError{
			Field:   "city",
			Message: "required",
		}
		fes = append(fes, err)
	}

	if req.Address.Country == "" {
		err := FieldError{
			Field:   "country",
			Message: "required",
		}
		fes = append(fes, err)
	}

	if len(fes) == 0 {
		return nil
	}

	return ValidationErrors{
		Errs: fes,
	}
}

func toEmployee(req EmployeeRequest) Employee {
	return Employee{
		Name:    req.Name,
		Age:     *req.Age,
		Address: req.Address,
		Skills:  req.Skills,
	}
}
