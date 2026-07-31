package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func WriteStringFunc(source io.Writer) {
	writer := bufio.NewWriter(source)
	defer writer.Flush()

	bytesWritten, err := writer.WriteString("Backend Engineering")
	if err != nil {
		fmt.Println("Error while writing into file", err)
		return
	}
	fmt.Println("bufio.WriteString()", bytesWritten)
}
