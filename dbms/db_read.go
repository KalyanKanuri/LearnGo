package dbms

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	UserName string
	Password string
	Email    string
}

func ReadFromUserDetails(db *sql.DB) {
	fmt.Println("-- Read From User Details --")
	var users []User
	query := `SELECT * FROM USER_DETAILS;`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}
	jsonResult, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonResult))
}
