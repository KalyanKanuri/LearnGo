package test

import (
	"io"
	"learning/go/test/pkg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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

// TestCheckUserNameWithAssert Using assert keyword from testify ext pkg
func TestCheckUserNameWithAssert(t *testing.T) {
	assert.Equal(t, false, pkg.CheckUserName("admin"))
}

// TestLogin
func TestLogin(t *testing.T) {
	got, err := pkg.Login("valid", "validPWD")
	assert.NoError(t, err)
	assert.Equal(t, true, got)
}

func TestLoginHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	pkg.LoginHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	body, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", string(body))
}
