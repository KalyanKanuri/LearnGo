package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func ReadBytesFunc(source io.Reader) {
	reader := bufio.NewReader(source)

	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("bufio.ReadBytes()", string(bytes))
			return
		}
		fmt.Println("Error reading bytes", err)
	}
	fmt.Println("bufio.ReadBytes()", string(bytes))
}
