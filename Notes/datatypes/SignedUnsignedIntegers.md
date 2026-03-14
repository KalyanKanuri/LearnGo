# Signed and Unsigned Integers in Go

Go provides both signed and unsigned integer types.

## Signed Integers

Signed integers can store both positive and negative values.

Common signed integer types:

- `int8`
- `int16`
- `int32`
- `int64`
- `int`

`int` is the general-purpose integer type in Go. Its size depends on the architecture:

- 32 bits on a 32-bit system
- 64 bits on a 64-bit system

Range formula for an `n`-bit signed integer:

`-2^(n-1)` to `2^(n-1) - 1`

Examples:

- `int8`: `-128` to `127`
- `int16`: `-32768` to `32767`
- `int32`: `-2147483648` to `2147483647`
- `int64`: `-9223372036854775808` to `9223372036854775807`

## Unsigned Integers

Unsigned integers can store only non-negative values.

Common unsigned integer types:

- `uint8`
- `uint16`
- `uint32`
- `uint64`
- `uint`

`uint` also depends on the architecture:

- 32 bits on a 32-bit system
- 64 bits on a 64-bit system

Range formula for an `n`-bit unsigned integer:

`0` to `2^n - 1`

Examples:

- `uint8`: `0` to `255`
- `uint16`: `0` to `65535`
- `uint32`: `0` to `4294967295`
- `uint64`: `0` to `18446744073709551615`

## How Go Stores Signed Integers

Go stores signed integers using **two's complement** representation, also written as **2's complement**.

This is the standard binary way to represent negative numbers in modern systems.

Important correction:

- It is not accurate to say "1 bit is used only for the sign"
- In two's complement, all bits participate in the value
- The leftmost bit still tells us whether the number is negative, but it is part of the numeric encoding, not a separate unused sign flag

## What Is Two's Complement?

Two's complement is a way to represent negative numbers in binary.

To get the negative version of a positive binary number:

1. Write the positive number in binary
2. Invert all bits
3. Add `1`

Example with `int8` and the value `5`:

```text
 5  = 00000101
invert bits:
      11111010
add 1:
     +00000001
      --------
-5  = 11111011
```

So in an 8-bit signed integer:

- `5` is `00000101`
- `-5` is `11111011`

## Why the Signed Range Looks Different

For `int8`, there are 8 bits total, so there are `2^8 = 256` possible bit patterns.

Those patterns are split like this:

- `128` negative values: `-128` to `-1`
- `1` zero value: `0`
- `127` positive values: `1` to `127`

That is why `int8` goes from `-128` to `127`, not `-127` to `127`.

The smallest value has no positive mirror:

```text
10000000 = -128
```

## Unsigned vs Signed Bit Meaning

For `uint8`, the bits represent:

```text
128 64 32 16 8 4 2 1
```

For `int8` in two's complement, the bits represent:

```text
-128 64 32 16 8 4 2 1
```

That is why the leftmost bit changes the meaning of the number.

## Simple Go Example

```go
package main

import "fmt"

func main() {
	var temperature int = -25
	var age uint = 20

	fmt.Println(temperature)
	fmt.Println(age)

	var a int8 = 5
	var b int8 = -5

	fmt.Printf("5  -> %08b\n", uint8(a))
	fmt.Printf("-5 -> %08b\n", uint8(b))
}
```

Output:

```text
5  -> 00000101
-5 -> 11111011
```

`uint8(b)` is used only to print the raw 8-bit pattern.

## Practical Rule

- Use `int` for most normal integer work in Go
- Use `uint` only when the value truly cannot be negative or when working with bit patterns
- Be careful when converting between signed and unsigned integers, because negative values can become large positive numbers
