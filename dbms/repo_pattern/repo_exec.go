package repopattern

import (
	"database/sql"
	"fmt"
	"learning/go/dbms/repo_pattern/models"
	"learning/go/dbms/repo_pattern/repo"
	"learning/go/dbms/repo_pattern/utils"

	_ "modernc.org/sqlite"
)

func connectToDb() *sql.DB {
	db, err := sql.Open("sqlite", "demo.db")
	repoutils.CheckErr(err)
	return db
}

func RepoPatternInGo() {
	fmt.Println("-- Repository Pattern in Go --")
	db := connectToDb()

	customerRepo := repo.NewCustomerRepo(db)
	newCustProfile := models.CustProfile{
		IsActive: true,
		Role:     "User",
	}
	newCustomer := models.Customer{
		Name:        "John",
		Address:     "Manhattan, Vice City, GTA",
		CustProfile: newCustProfile,
	}
	customerRepo.CreateCustomer(&newCustomer)
	john := customerRepo.GetCustomer("John")

	orderRepo := repo.NewOrderRepo(db)

	var items []models.OrderItem
	newOrdItem := models.OrderItem{
		ProductName: "MSI Gaming Laptop",
		ProductType: "Electronics",
		Quantity:    1,
		UnitPrice:   65000,
	}
	items = append(items, newOrdItem)
	newOrder := models.Order{
		Description: "Laptop Order",
		Price:       newOrdItem.UnitPrice * float64(newOrdItem.Quantity),
		Items:       items,
		Customer:    john,
	}
	orderID := orderRepo.CreateOrder(&newOrder)
	orderRepo.GetOrderDetails(orderID)
	orderRepo.GetCustomerOrders("John")
}
