package dbms

import (
	"database/sql"
	"fmt"
)

/*
 * Transactions in Go
 * ------------------
 * transactions has 3 parts to do
 * Begin
 * Optional Rollback if any error
 * Commit to save the txn
 */
func TxnInGo(db *sql.DB) {
	fmt.Println("-- Transactions in Go --")
	txn, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmnt, err := txn.Prepare(`Insert into User_Details (ID, Username, password, email) values(?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err)
	}
	defer stmnt.Close()
	pwd := hashPwd("jadaSJHfv")
	res, err := stmnt.Exec(3, "Jane", string(pwd), "jane@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	userId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	profStmnt, err := txn.Prepare(`Insert into USER_PROFILE (USER_ID, IS_ACTIVE) values(?, ?)`)
	if err != nil {
		fmt.Println(err)
	}
	defer profStmnt.Close()
	res, err = profStmnt.Exec(userId, true)
	if err != nil {
		txn.Rollback()
		fmt.Println(err)
	}

	err = txn.Commit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.LastInsertId())
}
