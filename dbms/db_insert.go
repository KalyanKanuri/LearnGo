package dbms

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPwd(pwd string) []byte {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return hashedPwd
}

func InsertToUserDetails(db *sql.DB) (sql.Result, error) {
	fmt.Println("-- Insert Records into User Details --")
	query := `INSERT INTO USER_DETAILS (ID, USERNAME, PASSWORD, EMAIL) VALUES (?, ?, ?, ?)`
	uname := "John"
	pwd := "Johnpwd"
	email := "john@email.com"
	hashedPwd := hashPwd(pwd)
	result, err := db.Exec(query, 1, uname, string(hashedPwd), email)
	return result, err
}
