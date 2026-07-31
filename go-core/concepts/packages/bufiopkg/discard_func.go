package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func DiscardFunc(source io.Reader) {
	reader := bufio.NewReader(source)
	discarded, err := reader.Discard(4)
	if err != nil {
		fmt.Println("Error while discarding", err)
		return
	}
	fmt.Println("bufio.Discard()", discarded)
}
