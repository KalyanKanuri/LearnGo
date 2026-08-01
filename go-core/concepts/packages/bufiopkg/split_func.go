package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func BufSplit(source io.Reader) {
	scanner := bufio.NewScanner(source)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println("scaner.Split()", scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error spliting tokens", scanner.Err())
		return
	}
}
