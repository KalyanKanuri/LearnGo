# Channels

Channel are used for sharing data between goroutines, we have to create a channel with make() function as below

```go
ch := make(chan string) // provide the datatype of channel as per data passed
```

once the channel is created we can use this to pass data from one goroutine to another as below

```go
func sendMsg(ch chan string) {
    msg := "Hello from channel with sendMsg goroutine"
    ch <- msg // we are inserting into channel i.e. sending
}

func main() {
    ch := make(chan string)

    go sendMsg(ch)

    msg := <-ch // we are receiving from channel
    fmt.Println(msg)
}
```

no need to pass channels as pointers as channels are already reference types so even though we modify the
channel the underlying object will be updated instead of creating a new copy

by default channels are unbuffered so sending and receiving data should happen at the same time (synchronous)
so the sender blocks the goroutine until the receiver is ready and receiver blocks the goroutine until the receiver is ready

this is nothing but the goroutine pauses until the sender/receiver is ready.

channels send each msg from sender to receiver if we try to receive multiple msgs at receiver end when the sender is sending only one it will result in a deadlock as the receiver goroutine will wait forever while the receiver goroutine is already completed it's execution so the channel will not give the .
