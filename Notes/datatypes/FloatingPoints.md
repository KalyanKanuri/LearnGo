# Floating-Point Numbers in Go

Go provides two floating-point types:

- `float32`
- `float64`

These are used to store numbers with a fractional part, such as `3.14`, `0.5`, and `-12.75`.

## `float32` vs `float64`

- `float32` uses 32 bits of memory
- `float64` uses 64 bits of memory
- `float64` gives more precision than `float32`

In practice, `float64` is used more often because it is more precise and works directly with most functions in Go's `math` package.

## Default Floating-Point Type in Go

Floating-point literals like `3.14` start as untyped numeric constants.

When Go needs to infer a concrete type, it usually becomes `float64`.

Example:

```go
temp := 3.14      // inferred as float64
var pi float64 = 3.14
var rate float32 = 3.14
```

If you want to reduce memory usage and lower precision is acceptable, you can choose `float32`.

## Why Floating-Point Values Can Be Inexact

Floating-point numbers are stored in binary, not decimal.

Some decimal numbers, such as `0.1` and `0.2`, cannot be represented exactly in binary. Because of that, small rounding errors can appear.

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println(0.1 + 0.2)
}
```

Output:

```text
0.30000000000000004
```

This is called a **floating-point precision issue**.

The important idea is:

- the math is not "wrong"
- the stored binary approximation is slightly different from the decimal value you expect

## Comparing Floating-Point Values

Because of precision limits, comparing floating-point numbers directly with `==` can be risky.

Instead, compare whether the difference is very small.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(0.1+0.2-0.3) < 1e-9) // true
}
```

Here:

- `math.Abs` gives the absolute difference
- `1e-9` means `0.000000001`
- if the difference is smaller than that limit, we treat the values as equal

## When to Be Careful

Floating-point numbers are great for:

- measurements
- scientific calculations
- percentages
- graphics

Be more careful when working with money or exact decimal values, because floating-point types can introduce tiny rounding errors.

For financial values, it is often better to store:

- integer units such as cents or paise
- or use a decimal library when exact decimal behavior is required

## Practical Rule

- Use `float64` by default
- Use `float32` only when you specifically want lower memory usage
- Do not expect most decimal fractions to be stored exactly
- For equality checks, compare with a small tolerance instead of using `==`
