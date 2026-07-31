package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func WriteFunc(source io.Writer) {
	writer := bufio.NewWriter(source)
	defer writer.Flush()

	content := []byte{'H', 'e', 'l'}
	bytesWritten, err := writer.Write(content)
	if err != nil {
		fmt.Println("Error writing into file", err)
	}

	fmt.Println("bufio.Write()", bytesWritten)
}
