package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func ReadFunc(bufread io.Reader) {
	reader := bufio.NewReader(bufread)

	buf := make([]byte, 5)
	bytesRead, err := reader.Read(buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println(string(buf[:bytesRead]))
			return
		}
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println(string(buf[:bytesRead]))
}
