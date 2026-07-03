package concepts

import "fmt"

func ArraysInGo() {
	// array declaration
	var arr1 [5]int

	// array value infusion
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	// The zero[default] value of an integer array is zero[]
	fmt.Println(arr1) // result - [1 2 3 0 0]

	//array declaration and initialization
	arr2 := [5]int{1, 2}
	fmt.Println(arr2) // result - [1 2 0 0 0]

	// string array and zero value
	strArray := [5]string{"Apple", "Boy"}
	// zero value is a white space for string arrays
	fmt.Println(strArray) // result - [Apple Boy   ]

	// rune array and zero value
	runeArray := [5]rune{'1', 'A'}
	// zero value is zero
	// result came in numbers because rune is an alias of int32 and the data we gave in the array
	// is stored as an unicode code character of the given value
	fmt.Println(runeArray)            // result - [49 65 0 0 0]
	fmt.Println(string(runeArray[:])) // this will convert the unicode to string

	// byte array and zero value
	byteArray := [5]byte{'H', 'e', '1'}
	// zero value is zero
	// result cam in numbers as the data is stored bytes the data we gave will be converted to bytes
	fmt.Println(byteArray)            // result - [72 101 49 0 0]
	fmt.Println(string(byteArray[:])) // this will convert the byte to string

	// 2d array
	// explanation - first [2] this tells the number of items the outer array contains
	// 			    - second [3] this tells the numbers of items each inner array item contains
	twoDArray := [2][3]int{
		{1, 2},
		{1, 2},
	}
	fmt.Println(twoDArray)
}
