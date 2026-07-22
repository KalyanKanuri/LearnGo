# Sync.WaitGroup

WaitGroup is just a counter that lets user decides how many goroutines are getting scheduled and reduce the counter when a goroutine is completed.
It will ensure the function where the WaitGroup.Wait() is called is blocked until the goroutines finishes.
the goroutine sends a signal using WaitGroup.Done() method which unblocks the func and continues processing

## Race Conditions

Race Condition is a behaviour where the output becomes unpredicatble due to concurrency.
For Concurrency we are dependent on go scheduler and OS scheduler as well we are unaware which goroutine runs at which time exactly.
This is the reason why the output becomes unpredicatble in a concurrent pattern.

lets assume a scenario where we are in a restaurant

we have ordered chicken biryani, there are two chefs in the kitchen chef A sees the order note and starts preparing the order without any coordination with the other chef in the same way chef B sees the order note and starts preparing the order, now we have an unexpected result of two chicken biryanis this is nothing but race condition

when two or more threads are assigned a task related to same data, if there's no communication between the threads we cannot predict the result and the order of the output in programming terms see the below example

```go
func Worker(i int) {
    fmt.Println(i)
}

func main() {
    for i := range 1000 {
        go Worker(i)
    }
}
```

output can be as below, but it's unpredictable

```bash
991
1
3
5
997
1000
10
20
11
```

this is classical goroutine behaviour in the same way lets say we want to update a value, here where the problem of Race Condition comes lets see the example below

```go
func Update(wg *sync.WaitGroup, arr *[]int, val int) {
    defer wg.Done()
    *arr = append(*arr, val)
}

func main() {
    arr := make([]int, 10)
    var wg *sync.WaitGroup

    for i := range 10 {
        wg.Add(1)
        go Update(wg, &arr, i)
    }

    wg.Wait()
    fmt.Println("Final slice length:", len(arr))
}
```

here two goroutines might try to update the same index at same time one gorotuine with value 1 and the other might be with value 5 here the ouput is unexpected if we are expected.
