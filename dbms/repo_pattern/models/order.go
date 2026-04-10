package models

type Order struct {
	OrderID     string
	Description string
	Price       float64
	Items       []OrderItem
	Customer    Customer
	DateTimeMixin
}

type OrderItem struct {
	ItemID      int
	ProductName string
	ProductType string
	Quantity    int
	UnitPrice   float64
}
