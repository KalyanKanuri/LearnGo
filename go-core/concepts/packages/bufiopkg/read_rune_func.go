package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func ReadRuneFunc(source io.Reader) {
	reader := bufio.NewReader(source)
	r, size, err := reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			fmt.Println("bufio.ReadRune()", string(r), size)
			return
		}
		fmt.Println("Error reading rune", err)
		return
	}
	fmt.Println("bufio.ReadRune()", string(r), size)
}
