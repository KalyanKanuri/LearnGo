package extensions

import (
	"bufio"
	"fmt"
	"os"
)

func BufferReaderInGo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", name)
}
