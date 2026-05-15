package test

import (
	"learning/go/test/pkg"
	"testing"
)

/*
TestCheckUserName
this is Table Driven Test we can assume the data in below struct as data from a table
*/
func TestCheckUserName(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"Admin Test", "admin", false},
		{"Dot Test", "user.", false},
		{"Empty Test", "", false},
		{"Clean User", "username", true},
	}

	for _, tc := range testCases {
		got := pkg.CheckUserName(tc.input)
		if got != tc.want {
			t.Errorf("Test %s, failed expected %t and got %t",
				tc.name,
				tc.want,
				got,
			)
			return
		}
	}
}
