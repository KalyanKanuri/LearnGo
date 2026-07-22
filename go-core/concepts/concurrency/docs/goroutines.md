# Goroutines

Goroutines are simply structs managed by Go Scheduler, go scheduler schedules these goroutines into OS threads,
this is the reason we can create so many goroutines and the system can still work efficiently without having any performance issues.

A goroutine can be initialized simply as below

```go
func main() {
    counter := 0

    go func() {
        counter++
    }()
}
```
