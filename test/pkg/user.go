package pkg

import "strings"

func CheckUserName(userName string) bool {
	if userName == "" ||
		strings.Contains(userName, "admin") ||
		strings.Contains(userName, ".") {
		return false
	}
	return true
}
