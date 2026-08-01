package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func BufBuffered(source io.Writer) {
	writer := bufio.NewWriter(source)
	usedBytes := writer.Buffered()
	fmt.Println("bufio.Buffered()", usedBytes)
}
