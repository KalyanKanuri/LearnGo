# Context pkg

Context pkg is used to communicate with goroutines in a single channel where a broadcast message is published and cacaded to all goroutines within the context

## Use Cases

Suppose we have a user using our website and clicked on My Orders in his profile this will call the backend api and backend api will trigger the db call loading json for the response etc now if the user closes the browser in the middle of the request, do we need to still continue the process happening in backend? no to handle this context is introduced.

We give a context to all the goroutines triggered and can able to send a broadcast message to all the goroutines to signal that the context is cancelled and no need to continue anymore

Context is accessible with context pkg

```go
package context

// inside func
// creating a context with cancel
ctx, cancel := context.WithCancel(context.Background())


// in select we can use something as below
case <-ctx.Done()

// when cancel signal is sent the goroutine will be exited
cancel()
```

```go
// Context with timeout
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
// this context will automatically cancels after 1 second
// this takes relative time for e.g., 3 seconds
```

```go
ctx, cancel := context.WithDeadline(context.Background(), time.Add(time.Second))
defer cancel()
// this context will also automatically cancels after 1 second
// the only difference is this take absolute time for e.g., 2026-07-22 22:29:00
```
