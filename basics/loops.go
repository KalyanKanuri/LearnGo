package basics

import "fmt"

func LoopsInGo() {
	fmt.Println("\n******* Loops In Go *******")

	// For is the only loop in Go, it can be used in different ways

	// Tranditional for loop
	fmt.Println("-- Tranditional for loop --")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For-each loop
	fmt.Println("-- For-each loop --")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// While-like loop
	fmt.Println("-- While-like loop --")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop
	k := 5
	fmt.Println("-- Infinite loop --")
	for {
		fmt.Println("This will run forever")
		k++
		if k > 10 {
			fmt.Println("Breaking to avoid forever loop")
			break // Exit the loop when k is greater than 10
		}
	}
}
