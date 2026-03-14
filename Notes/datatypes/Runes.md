# Runes in Go

In Go, a `rune` is used to represent a single Unicode code point.

`rune` is just an alias for `int32`.

That means this is true:

```go
var r rune = 'A'
```

and conceptually it is the same kind of storage as:

```go
var r int32 = 65
```

## What a Rune Represents

A rune does not store a whole string.

A rune stores one Unicode code point, such as:

- `'A'`
- `'9'`
- `'@'`
- `'é'`
- `'😊'`

Example:

```go
var r rune = '😊'
```

Here, `r` stores the Unicode code point for that symbol.

## Unicode Code Point

Unicode gives a numeric code to each character or symbol.

For example:

- `'A'` is `U+0041`
- `'é'` is `U+00E9`
- `'😊'` is `U+1F60A`

The `U+...` form is hexadecimal notation for the Unicode code point.

So when you write:

```go
r := '😊'
```

Go stores the Unicode code point value for `😊` inside a `rune`.

For `😊`:

- Unicode code point: `U+1F60A`
- decimal value: `128522`

So yes, a rune ultimately stores an integer value, but the important meaning of that integer is: it represents a Unicode code point.

## Rune Literals

Rune literals use single quotes:

```go
'A'
'é'
'😊'
```

This is different from strings, which use double quotes:

```go
"A"
"hello"
"😊"
```

## Rune vs String

- a `rune` stores one Unicode code point
- a `string` stores a sequence of bytes

Example:

```go
var ch rune = 'G'
var text string = "Go"
```

Here:

- `ch` holds one character code point
- `text` holds multiple bytes that form a string

## Rune vs Byte

This is an important difference in Go:

- `byte` is an alias for `uint8`
- `rune` is an alias for `int32`

Usually:

- `byte` is used for raw data or UTF-8 bytes
- `rune` is used when you want to work with Unicode characters

## Simple Example

```go
package main

import "fmt"

func main() {
	r := '😊'

	fmt.Printf("%T\n", r)
	fmt.Printf("%U\n", r)
	fmt.Printf("%d\n", r)
}
```

Possible output:

```text
int32
U+1F60A
128522
```

This shows:

- the underlying type is `int32`
- the Unicode code point is `U+1F60A`
- the decimal integer value is `128522`

## One Important Note

A rune represents one Unicode code point, not always one "visual character".

Some visible characters can be formed from multiple code points, so "one character" and "one rune" are not always exactly the same idea.

## Practical Rule

- Use `rune` when working with Unicode characters
- Use `string` for text
- Use `byte` when working with raw bytes or UTF-8 encoded data
