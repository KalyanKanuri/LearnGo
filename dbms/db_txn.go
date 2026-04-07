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

}
