package bufiopkg

import (
	"bufio"
	"fmt"
	"io"
)

func WriteRuneFunc(source io.Writer) {
	writer := bufio.NewWriter(source)
	defer writer.Flush()

	runeSize, err := writer.WriteRune('😂')
	if err != nil {
		fmt.Println("Error writing rune", err)
		return
	}
	fmt.Println("bufio.WriteRune()", runeSize)
}
