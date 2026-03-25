package extensions

import (
	"bufio"
	"fmt"
	"os"
)

func BufferReaderInGo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name: ")
	name, _, _ := reader.ReadLine()
	fmt.Printf("Hello %s", name)
}
