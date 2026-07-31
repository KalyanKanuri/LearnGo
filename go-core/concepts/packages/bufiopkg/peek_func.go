package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func PeekFunc(source io.Reader) {
	reader := bufio.NewReader(source)

	bytes, err := reader.Peek(4)
	if err != nil {
		fmt.Println("Error peeking bytes", err)
		return
	}
	fmt.Println("bufio.Peek()", string(bytes))
}
