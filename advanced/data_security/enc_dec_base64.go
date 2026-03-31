package datasecurity

import (
	"encoding/base64"
	"fmt"
)

func Base64EncodeDecodeInGo() {
	fmt.Println("-- Base64 Encoding in Go --")
	bs64String := "Some Secure data need to be encoded"
	encodedString := base64.StdEncoding.EncodeToString([]byte(bs64String))
	fmt.Println(encodedString)

	decodedString, _ := base64.StdEncoding.DecodeString(encodedString)
	fmt.Println(string(decodedString))
}