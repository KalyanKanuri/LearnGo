package core

import (
	"fmt"
)

type User struct {
	ID     int
	Name   string
	Access bool
}

// we have embedded User struct into Admin struct, Go promotes the values of User to Admin
// i.e., Admin also contains ID and Name now. this concept is called Embedding
type Admin struct {
	*User
	Group string
	Level int
}

type AccessProvider interface {
	provideAccess()
}

func (u *User) provideAccess() {
	fmt.Println("Process Access privileges")
	u.Access = true
	fmt.Println("Access granted", u.Access)
}

func EmbeddingInGo() {
	fmt.Println("\n******* Embedding In Go *******")
	joe := &User{
		ID:     123,
		Name:   "Joe",
		Access: false,
	}

	salesAdmin := Admin{
		User:  joe,
		Group: "Sales",
		Level: 1,
	}

	fmt.Printf("salesAdmin: user --> %+v, admin --> group - %s, level - %d\n",
		*salesAdmin.User,
		salesAdmin.Group,
		salesAdmin.Level,
	)
}
