package basics

import "fmt"

func ConditionalsInGo() {
	fmt.Println("\n******* Conditionals In Go *******")
	canCode := true
	hasExperience := false
	var role string

	fmt.Println("-- If-else_if-else statement --")
	if canCode && hasExperience {
		role = "Senior Developer"
		fmt.Println("Assign to Coding profession as Senior Developer")
	} else if canCode && !hasExperience {
		role = "Junior Developer"
		fmt.Println("Assign to Coding profession as Junior Developer")
	} else if !canCode && hasExperience {
		role = "Senior Business Analyst"
		fmt.Println("Assign to non-coding profession to a Senior role")
	} else {
		role = "Junior Business Analyst"
		fmt.Println("Assign to non-coding profession to a Junior role")
	}

	fmt.Println("-- Switch statement --")
	switch role {
	case "Senior Developer":
		fmt.Println("No Coding Training required can be assigned to Project")
	case "Junior Developer":
		fmt.Println("Coding Training required, can't be assigned to Project yet")
	case "Senior Business Analyst":
		fmt.Println("No Business Analysis Training required can be assigned to Project")
	case "Junior Business Analyst":
		fmt.Println("Business Analysis Training required, can't be assigned to Project yet")
	}
}
