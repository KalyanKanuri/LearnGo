package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func WriteByteFunc(source io.Writer) {
	writer := bufio.NewWriter(source)
	defer writer.Flush()

	err := writer.WriteByte('h')
	if err != nil {
		fmt.Println("Error writing byte", err)
		return
	}
	fmt.Println("bufio.WriteByte() - returns nothing")
}
