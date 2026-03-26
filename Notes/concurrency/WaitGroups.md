# WaitGroups

In Go when we want to control over go routines to not to halt without executing when the main goroutine
execution is completed, we can use time. Sleep() by providing how much time to sleep as per the goroutine
execution time, but we cannot go on maintaining time to sleep when the goroutine execution time is increased
we have to increase the main goroutine sleep time as well, this can be a hectic process.

To encounter this to process main goroutine wait dynamically we have WaitGroups we can give the number of
goroutines as below

```go
var wg sync.WaitGroup.Add()
wg.Add(4) // 4 represents the number of goroutines
```

## Rules to Follow

* Always use wg. Add() outside the goroutine which we are trying to invoke
* Must use wg. Done() inside the goroutine which is invoked to decrease the added goroutine counter
* Always use wg. Wait() to wait for the invoked goroutines execution completion
* Must pass *sync. WaitGroup pointer to the invoked pointer to avoid creating wg copy
