package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func ReadStringFunc(source io.Reader) {
	reader := bufio.NewReader(source)

	str, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("bufio.ReadString()", str)
			return
		}
		fmt.Println("Error while reading string", err)
		return
	}
	fmt.Println("bufio.ReadString()", str)
}
