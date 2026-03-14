# Complex Numbers in Go

Go supports complex numbers as built-in numeric types.

A complex number has:

- a real part
- an imaginary part

A value like `3 + 4i` means:

- real part: `3`
- imaginary part: `4`

## Complex Types in Go

Go provides two complex types:

- `complex64`
- `complex128`

They are based on floating-point values:

- `complex64` = `float32` real part + `float32` imaginary part
- `complex128` = `float64` real part + `float64` imaginary part

In practice, `complex128` is used more often because it gives more precision.

## Declaring Complex Numbers

You can write complex numbers directly using `i`.

```go
var x complex64 = 1 + 2i
var y complex128 = 3 + 4i
```

You can also create them with the built-in `complex` function:

```go
var x complex64 = complex(1, 2)
var y complex128 = complex(3, 4)
```

The `complex(real, imag)` function builds a complex value from two numbers.

## Default Complex Type

When Go infers the type from a complex literal, it usually becomes `complex128`.

Example:

```go
z := 2 + 3i // inferred as complex128
```

This matches the usual Go behavior of preferring `float64`-based values when a concrete type is inferred.

## Operations on Complex Numbers

You can use normal arithmetic operators with complex values:

- addition: `+`
- subtraction: `-`
- multiplication: `*`
- division: `/`

Example:

```go
package main

import "fmt"

func main() {
	x := 1 + 2i
	y := 3 + 4i

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
}
```

## Getting the Real and Imaginary Parts

Go provides two built-in functions:

- `real()`
- `imag()`

Example:

```go
var x complex64 = 1 + 2i

var realPart float32 = real(x)
var imaginaryPart float32 = imag(x)
```

For `complex64`, `real(x)` and `imag(x)` return `float32`.

For `complex128`, they return `float64`.

## Simple Example

```go
package main

import "fmt"

func main() {
	z := 5 + 7i

	fmt.Println(z)
	fmt.Println(real(z))
	fmt.Println(imag(z))
}
```

Possible output:

```text
(5+7i)
5
7
```

## Practical Rule

- Use `complex128` by default
- Use `complex64` only when lower memory usage is important
- Use `real()` and `imag()` when you need the two parts separately
- Complex numbers are useful in math, engineering, graphics, and signal processing
