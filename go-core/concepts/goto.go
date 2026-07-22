package concepts

import (
	"fmt"
)

func GoToStatement() {
	// goto statement is used to transfer control to a labeled statement in the same function.
	// it is generally used to break out of nested loops or to jump to a specific point in the code.
	// however, it is not recommended to use goto statement as it can make the code less readable and harder to maintain.
	i := 0
LoopStart:
	if i >= 3 {
		goto LoopEnd
	}
	i++
	fmt.Println("i:", i)
	goto LoopStart
LoopEnd:
	fmt.Println("Done")
}
