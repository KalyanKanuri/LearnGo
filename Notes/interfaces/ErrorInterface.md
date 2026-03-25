# Error Interface in Go

In Go, `error` is a built-in interface used to represent something that went wrong.

Its definition is:

```go
type error interface {
	Error() string
}
```

This means any type that has a method:

```go
Error() string
```

automatically satisfies the `error` interface.

## Why Go Uses an Interface for Errors

Go does not use exceptions for normal error handling.

Instead, functions usually return an `error` value along with their main result.

Example:

```go
file, err := os.Open("data.txt")
if err != nil {
	fmt.Println("could not open file:", err)
	return
}
```

Here:

- `file` is the successful result
- `err` tells us whether something went wrong

If `err` is `nil`, the operation succeeded.

## The Zero Value of an Error

The zero value of an `error` is `nil`.

That means:

```go
var err error
fmt.Println(err == nil)
```

Output:

```text
true
```

So in Go:

- `nil` means no error
- non-`nil` means an error happened

## Common Function Pattern

A very common Go pattern is:

```go
value, err := someFunction()
if err != nil {
	return err
}
```

This keeps error handling explicit and easy to follow.

## Creating Errors

One common way to create an error is with `errors.New`.

```go
import "errors"

err := errors.New("something went wrong")
fmt.Println(err)
```

This creates a simple error value whose message comes from the string.

Another common way is `fmt.Errorf`.

```go
err := fmt.Errorf("user %d not found", 42)
```

This is useful when you want formatted error messages.

## Custom Error Types

Because `error` is an interface, you can create your own error types.

Example:

```go
package main

import "fmt"

type ValidationError struct {
	Field string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s is invalid", v.Field)
}

func main() {
	var err error = ValidationError{Field: "email"}
	fmt.Println(err)
}
```

Since `ValidationError` has an `Error() string` method, it satisfies the `error` interface.

## Error Interface in Plain Words

You can think of the `error` interface like this:

- Go does not care what the concrete error type is
- Go only cares that it can call `Error()` on it

That is the power of interfaces in Go.

Different error types can exist, but all of them can be treated as `error`.

## Wrapping Errors

Sometimes one function wants to add more context to an error it received.

Example:

```go
data, err := os.ReadFile("config.json")
if err != nil {
	return fmt.Errorf("reading config file: %w", err)
}
```

Here `%w` wraps the original error so it can still be checked later.

This is better than losing the original error information.

## Checking Errors

Basic check:

```go
if err != nil {
	fmt.Println("error:", err)
}
```

For wrapped or specific errors, Go provides `errors.Is` and `errors.As`.

Example with `errors.Is`:

```go
if errors.Is(err, os.ErrNotExist) {
	fmt.Println("file does not exist")
}
```

Example with `errors.As`:

```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
	fmt.Println("path error happened on:", pathErr.Path)
}
```

## Interface Satisfaction Example

This type also becomes an error:

```go
type MyError struct{}

func (m MyError) Error() string {
	return "my custom error"
}
```

Then:

```go
var err error = MyError{}
fmt.Println(err.Error())
```

This works because `MyError` satisfies the `error` interface.

## Important Practical Rule

If a function returns an `error`, always check it unless you have a very specific reason not to.

For example, this ignores the error:

```go
name, _ := reader.ReadString('\n')
```

This works for quick demos, but in real programs it is better to handle the error:

```go
name, err := reader.ReadString('\n')
if err != nil {
	fmt.Println("failed to read input:", err)
	return
}
```

## One Important Subtlety

Be careful with typed `nil` values stored inside an interface.

Example:

```go
var e *MyError = nil
var err error = e

fmt.Println(err == nil)
```

Output:

```text
false
```

Even though `e` is `nil`, `err` is not `nil` because the interface still holds type information.

This is a common beginner gotcha in Go.

## Practical Summary

- `error` is a built-in interface
- its method is `Error() string`
- `nil` means no error
- non-`nil` means something failed
- `errors.New` creates simple errors
- `fmt.Errorf` creates formatted errors
- custom types can satisfy the `error` interface
- use `%w` to wrap errors
- use `errors.Is` and `errors.As` to inspect errors
- do not ignore errors in real programs unless it is intentional
