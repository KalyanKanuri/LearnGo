# Buffer Reader in Go

In Go, `bufio.NewReader` is commonly used when we want to read input from the console.

It is especially useful when:

- we want to read a full line of input
- input may contain spaces
- we want more control than `fmt.Scan` or `fmt.Scanln`

## Basic Idea

A buffered reader reads data through a buffer instead of reading one small piece directly every time.

In simple words:

- `os.Stdin` gives access to standard input
- `bufio.NewReader(os.Stdin)` wraps that input source
- the reader helps us read strings, bytes, runes, or lines more conveniently

## Creating a Reader for Console Input

Example:

```go
reader := bufio.NewReader(os.Stdin)
```

Here:

- `bufio` is the package for buffered I/O
- `os.Stdin` is the keyboard input stream
- `reader` is used to read what the user types

## Simple Console Input Example

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read input:", err)
		return
	}

	fmt.Printf("Hello %s", name)
}
```

This reads input until the user presses Enter.

## What `ReadString('\n')` Means

This line:

```go
name, err := reader.ReadString('\n')
```

means:

- read characters from the input
- keep reading until newline `'\n'` is found
- return the full text as a string

If the user types:

```text
Kalyan
```

then `name` usually becomes:

```text
"Kalyan\n"
```

Notice that the newline is included in the returned string.

## Removing the Newline

Usually we do not want the trailing newline in the final value.

So we often use `strings.TrimSpace`:

```go
name = strings.TrimSpace(name)
```

Example:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read input:", err)
		return
	}

	name = strings.TrimSpace(name)
	fmt.Printf("Hello %s\n", name)
}
```

Now if the user types `Kalyan`, the output is cleaner.

## Why Not Just Use `fmt.Scanln`

`fmt.Scanln` is simpler for basic input, but it stops at spaces.

For example, if the user enters:

```text
John Doe
```

`fmt.Scanln` may only read part of it depending on how it is used.

Buffered readers are more flexible for full-line input.

## Reading Input with Spaces

This is one of the biggest reasons to use a buffered reader.

Example:

```go
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter your full name: ")
fullName, err := reader.ReadString('\n')
if err != nil {
	fmt.Println("failed to read input:", err)
	return
}

fullName = strings.TrimSpace(fullName)
fmt.Println("Full name:", fullName)
```

This works well for input like:

```text
John Doe
```

## Reading Numbers from Console

Buffered readers return text, so if you want a number, you usually convert the string.

Example:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your age: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read input:", err)
		return
	}

	input = strings.TrimSpace(input)

	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("invalid number:", err)
		return
	}

	fmt.Println("Age:", age)
}
```

So the flow is:

- read text from console
- remove spaces or newline
- convert to the required type

## Other Useful Reader Methods

`bufio.Reader` has several methods.

Some common ones are:

- `ReadString(delim)` returns a string until a delimiter
- `ReadBytes(delim)` returns bytes until a delimiter
- `ReadRune()` reads one rune at a time
- `ReadByte()` reads one byte at a time

For console input, `ReadString('\n')` is one of the most common choices.

## What Happens on EOF

If input ends before the delimiter is found, the function may return:

- the data read so far
- an error such as `io.EOF`

This matters more when reading from files or redirected input than in simple keyboard demos.

## Important Practical Rule

Do not ignore the returned error in real programs.

For example, this is not ideal:

```go
name, _ := reader.ReadString('\n')
```

It is better to write:

```go
name, err := reader.ReadString('\n')
if err != nil {
	fmt.Println("failed to read input:", err)
	return
}
```

## Buffer Reader in Plain Words

You can think of it like this:

- `os.Stdin` is the source of keyboard input
- `bufio.NewReader` makes reading easier
- `ReadString('\n')` reads until Enter is pressed
- `strings.TrimSpace` cleans the final input

## When to Use It

Use `bufio.NewReader(os.Stdin)` when:

- you want full-line input
- input may contain spaces
- you want to read text comfortably from the console

## Practical Summary

- use `bufio.NewReader(os.Stdin)` for console input
- `ReadString('\n')` reads until newline
- the returned string often includes `\n`
- use `strings.TrimSpace` to clean input
- check the returned `error`
- convert the string later if you need numbers or other types
