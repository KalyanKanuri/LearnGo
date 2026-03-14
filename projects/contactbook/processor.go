package contactbook

import (
	"fmt"
)

type Contact struct {
	Name  string
	Age   int
	Email string
	Phone string
}

var ContactBook = make(map[string]*Contact)

func init() {
	fmt.Println("Initializing contact book...")
	ContactBook = make(map[string]*Contact)
}

func checkContactExists(name string) bool {
	_, exists := ContactBook[name]
	return exists
}

func GetContact(name string) (Contact, bool) {
	contact, exists := ContactBook[name]
	return *contact, exists
}

func AddContact(contact *Contact) {
	exists := checkContactExists(contact.Name)
	if exists {
		fmt.Printf("Contact with name '%s' already exists.\n", contact.Name)
		return
	}
	ContactBook[contact.Name] = contact
	fmt.Printf("Contact details for '%s' added successfully.\n", contact.Name)
}

func ListContacts() {
	fmt.Println(ContactBook)
	// In structs go automatically dereferences the pointer
	// when accessing the fields, so we can directly access the fields without using *contact.
	// only when we want to modify the contact or insert a new contact, we need to use the pointer.
	for id, contact := range ContactBook {
		fmt.Printf("ID: %s, Name: %s, Age: %d, Email: %s, Phone: %s\n",
			id,
			contact.Name,
			contact.Age,
			contact.Email,
			contact.Phone,
		)
	}
}
