package core

import (
	"fmt"
	"time"
)

/*
 * OOPS In Go
 * in Go we truly don't do OOPS programming but
 * we can achieve the OOPS pattern like programming
 * with structs, methods, recievers, interfaces etc.,
 * we can mimic OOP but it's not true OOP
 */

// this acts like a class but without those special abilities like in other languages.
type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	IsActive   bool
	Department string
	JoinedAt   time.Time
}

// this is like a constructor for Employee struct
func NewEmployee(
	id int, firstName, lastName, email, phone string, isActive bool, department string,
) Employee {
	employee := Employee{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Phone:      phone,
		IsActive:   isActive,
		Department: department,
		JoinedAt:   time.Now(),
	}
	return employee
}

/*
 * the reason we are using pointer over here is go does not refer
 * to the actual object passed to it, it creates a copy of the object
 * so when we modify anything we are just modifying the copy of the
 * object so we have to use the pointer to refer to the actual object
 * which is passed for updating the fields of the object
 *
 * here (e *Employee) is called a receiver -> pointer receiver
 * using receivers we can mimic OOPS like programming in Go
 * we have value receivers as well, we have to keep in mind
 * we have to use value receivers for only getters type of funcs
 * or only for getting date not modifying or inserting date
 */
func (e *Employee) UpdateActiveState() {
	e.IsActive = !e.IsActive
}

// here as we are only trying to get data we can use value receivers as well
func (e *Employee) GetEmployeeFullName() string {
	fullName := fmt.Sprintf("%+v, %+v", e.LastName, e.FirstName)
	return fullName
}
