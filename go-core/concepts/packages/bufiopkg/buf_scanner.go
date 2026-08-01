package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

// Buffer Scanner has a token limit of 64 KB, it processes 64 KB line at once if the line is larger it would error saying token size too long
func BufScanner(source io.Reader) {
	scanner := bufio.NewScanner(source)

	for scanner.Scan() {
		fmt.Println("bufio.Scanner", scanner.Text())
		// we can use scanner.Bytes() as well to get the raw bytes instead of string
	}

	if scanner.Err() != nil {
		fmt.Println("Error Scanning reader", scanner.Err())
	}
}
