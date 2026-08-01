package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func BufAvailable(source io.Writer) {
	writer := bufio.NewWriter(source)
	unusedBytes := writer.Available()
	fmt.Println("bufio.Available()", unusedBytes)
}
