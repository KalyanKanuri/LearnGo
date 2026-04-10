package models

type Customer struct {
	CustomerID int
	Name       string
	Address    string
	Orders     []Order
	CustProfile
	DateTimeMixin
}

type CustProfile struct {
	IsActive bool
	Role     string
}
