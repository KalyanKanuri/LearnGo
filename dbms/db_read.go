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
	IsActive bool
}

func ReadFromUserDetails(db *sql.DB) {
	fmt.Println("-- Read From User Details --")
	var users []User
	query := `SELECT UD.USER_ID, UD.USERNAME, UD.PASSWORD, UD.EMAIL, UP.IS_ACTIVE FROM USER_DETAILS UD JOIN USER_PROFILE UP ON UD.USER_ID = UP.USER_ID;`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Email, &user.IsActive)
		users = append(users, user)
	}
	jsonResult, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonResult))
}
