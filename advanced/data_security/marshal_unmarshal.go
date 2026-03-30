package datasecurity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	ID       int    `json:"id" xml:"ID"`
	Name     string `json:"username"  xml:"UserName"`
	Contact  string `json:"phone_no"  xml:"PhoneNo"`
	IsActive bool   `json:"active"  xml:"Active"`
}

var jsonPayload = `
{
 "id": 1,
 "username": "Jane",
 "phone_no": "12345678",
 "active": true
}
`

var xmlPayload = `
<User>
 <ID>1</ID>
 <Name>Jane</Name>
 <Contact>12345678</Contact>
 <IsActive>true</IsActive>
</User>
`

func MarshalUnMarshalInGo() {
	fmt.Println("-- Marshal Unmarshal in Go --")
	Jane := &User{
		ID:       1,
		Name:     "Jane",
		Contact:  "12345678",
		IsActive: true,
	}

	// we can use json.MarshalIndent for better visibility
	// but at computational standpoint json.Marshal is much cheaper
	jsonBytes, err := json.MarshalIndent(*Jane, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json marshal:\n%+v\n", string(jsonBytes))

	xmlBytes, err := xml.MarshalIndent(*Jane, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("xml marshal:\n%+v\n", string(xmlBytes))

	var jsonU User
	err = json.Unmarshal([]byte(jsonBytes), &jsonU)
	fmt.Println("After JSON Unmarshal", jsonU)

	var xmlU User
	err = xml.Unmarshal([]byte(xmlBytes), &xmlU)
	fmt.Println("After XML Unmarshal", xmlU)
}
