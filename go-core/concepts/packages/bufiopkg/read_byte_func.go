package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func ReadByteFunc(source io.Reader) {
	// reads only one byte at a time
	reader := bufio.NewReader(source)
	byt, err := reader.ReadByte()
	if err != nil {
		if err == io.EOF {
			fmt.Println("bufio.ReadByte()", string(byt))
			return
		}
		fmt.Println("Error while reading byte", err)
		return
	}
	fmt.Println("bufio.ReadByte()", string(byt))
}
