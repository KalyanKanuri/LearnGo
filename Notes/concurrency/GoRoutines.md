# GoRoutines

go routines are described as light weight threads in go, these are controlled
by goruntime instead of OS.

* when the main goroutine is completed it's execution the other goroutines will be halted in the middle and exists execution

we can start a goroutine as below

```go

func task() {
    // some task processing code
}

func main() {
    go task() // simply we can use go keyword to start a goroutine
}
```
