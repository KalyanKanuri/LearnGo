package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func CheckUserName(userName string) bool {
	if userName == "" ||
		strings.Contains(userName, "admin") ||
		strings.Contains(userName, ".") {
		return false
	}
	return true
}

func Login(userName, pwd string) (bool, error) {
	if userName == "" && pwd == "" {
		return false, errors.New("Empty Username and password provided")
	}
	if !CheckUserName(userName) {
		return false, fmt.Errorf("Inavlid Username %s", userName)
	}

	if userName == "valid" && pwd == "validPWD" {
		return true, nil
	}
	return true, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
