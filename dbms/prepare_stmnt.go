package dbms

import (
	"database/sql"
	"fmt"
)

func PrepareStmntInGo(db *sql.DB) {
	fmt.Println("-- Prepare Statement in Go --")
	stmnt, err := db.Prepare(`Insert into User_Details (ID, Username, password, email) values(?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmnt.Exec(2, "Jo", "joskjf", "jo@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.LastInsertId())
}