# Sync.Mutex [Mutual Exclusion]

Mutex ensures that a critical section is accessed by goroutines without causing any overlap, this causes the goroutines to never have a race condition

## Critical Section

A critical section is simply the section of code which needs protection from concurrency the other parts of the code can run concurrently, for e.g., the coed block where we accesses shared data

```go
func main() {
    var wg *sync.WaitGroup
    var mu sync.Mutex

    counter := 0

    for i := range 5 {
        wg.Add(1)

        go func() {
            defer wg.Done()
            defer mu.Unlock()

            mu.Lock()
            counter++
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
```
