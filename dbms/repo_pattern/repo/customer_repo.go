package repo

import (
	"database/sql"
	"fmt"
	"learning/go/dbms/repo_pattern/models"
	"learning/go/dbms/repo_pattern/utils"
)

type CustomerRepo interface {
	CreateCustomer(cust *models.Customer)
	GetCustomer(name string) models.Customer
}

type CustRepo struct {
	DB *sql.DB
}

func NewCustomerRepo(db *sql.DB) *CustRepo {
	return &CustRepo{
		DB: db,
	}
}

func (cr *CustRepo) CreateCustomer(cust *models.Customer) int {
	fmt.Printf("[CreateCustomer] Creating customer: %+v\n", cust)
	txn, err := cr.DB.Begin()
	repoutils.CheckErr(err)
	defer txn.Rollback()
	fmt.Println("[CreateCustomer] Transaction started")

	customerSql := `
				INSERT INTO customer (customer_name, customer_address)
				VALUES (?, ?)
			`
	stmnt, err := txn.Prepare(customerSql)
	repoutils.CheckErr(err)
	defer stmnt.Close()
	custRes, err := stmnt.Exec(cust.Name, cust.Address)
	repoutils.CheckErr(err)

	custID, err := custRes.LastInsertId()
	repoutils.CheckErr(err)
	fmt.Printf(
		"[CreateCustomer] Customer inserted | ID=%d Name=%s Address=%s\n",
		custID, cust.Name, cust.Address,
	)

	custProfileSql := `
					INSERT INTO customer_profile (customer_id, is_active, role)
					VALUES (?, ?, ?)
				`
	stmnt, err = txn.Prepare(custProfileSql)
	repoutils.CheckErr(err)
	defer stmnt.Close()
	custProfRes, err := stmnt.Exec(custID, cust.IsActive, cust.Role)
	repoutils.CheckErr(err)

	custProfID, err := custProfRes.LastInsertId()
	repoutils.CheckErr(err)
	fmt.Printf(
		"[CreateCustomer] Customer profile inserted | ProfileID=%d IsActive=%v Role=%s\n",
		custProfID, cust.IsActive, cust.Role,
	)

	err = txn.Commit()
	repoutils.CheckErr(err)
	fmt.Println("[CreateCustomer] Transaction committed successfully")
	fmt.Printf("[CreateCustomer] Created CustomerProfile ID=%d\n", custProfID)
	return cust.CustomerID
}

func (cr CustRepo) GetCustomer(name string) models.Customer {
	var customer models.Customer

	fmt.Printf("[GetCustomer] Fetching customer with name=%s\n", name)
	selectCustSql := `
					SELECT
						c.customer_id,
						c.customer_name,
						c.customer_address,
						cp.is_active,
						cp.role,
						c.created_at
					FROM customer c JOIN customer_profile cp
						ON c.customer_id = cp.customer_id
					WHERE c.customer_name = ?
				`
	stmnt, err := cr.DB.Prepare(selectCustSql)
	repoutils.CheckErr(err)
	defer stmnt.Close()

	cust := stmnt.QueryRow(name)
	err = cust.Scan(
		&customer.CustomerID,
		&customer.Name,
		&customer.Address,
		&customer.IsActive,
		&customer.Role,
		&customer.CreatedAt,
	)
	repoutils.CheckErr(err)
	fmt.Printf(
		"[GetCustomer] Customer found | ID=%d Name=%s Active=%v Role=%s\n",
		customer.CustomerID,
		customer.Name,
		customer.IsActive,
		customer.Role,
	)
	return customer
}
