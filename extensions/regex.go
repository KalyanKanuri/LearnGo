package extensions

import (
	"fmt"
	"regexp"
)

func RegexInGo() {
	fmt.Println("-- Using regexp.Compile --")
	text1 := "Learning regex in Go"
	compiled, err := regexp.Compile(`Go`)
	if err != nil {
		fmt.Println(err)
	} else {
		result := compiled.Match([]byte(text1))
		fmt.Printf("Processing regex using pattern: %+v, on text: %s, match: %t\n", compiled, text1, result)
	}
	fmt.Println("-- Using regexp.MustCompile --")
	text2 := "Product Codes: P-123, P-125, Serv-123"
	compiled = regexp.MustCompile(`P-\d+`)
	result := compiled.FindAllString(text2, -1)
	fmt.Printf("product number regex --> pattern: %+v, text: %+v result: %+v\n", compiled, text2, result)
	email := "test@gmail.com"
	pattern := `[a-zA-z]+@[a-zA-z]+\.(com|net)`
	compiled = regexp.MustCompile(pattern)
	matchResult := compiled.MatchString(email)
	fmt.Printf("Email validator result --> pattern: %+v, text: %+v, result: %t\n", pattern, email, matchResult)
}
