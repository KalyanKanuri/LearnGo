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

## Unbuffered vs Buffered channels

buffer means a temporary storage we have already discussed about the sender and receiver pattern used by channels this usually blocks the sender goroutine and receiver goroutine if we want to avoid that we can leverage buffered channels which can create with make by allocating certain memory for it as below

```go
func sender(ch chan string) {
    ch <- "sending msg from sender() goroutine"
}

func main() {
    // we have allocated 5 blocks of memory for the channel to use for temporary hold of msg
    ch := make(chan, string, 5)

    go sender(ch)

    // before the msg is received from sender here it will be placed to channel buffer
    // when the channel buffer is full it goes back to synchronous mode which is the usual working
    // pattern of unbuffered channel
    msg := <-ch

    fmt.Println(msg)
}
```

## Closing channels

channels are like pipelines which allows dataflow when the channel is not closed there might be a possibility of memory leaks as same as water wastage when a tap is left open for a water pipeline so to achieve it channels have something called close() to close the channel we can just the channel object to the func and the channel will be closed as below.

```go
func main() {
    ch := make(chan, string)
    go func() {
        ch <- "some message"
    }()
    msg := <-ch
    close(ch) // closing a channel
}
```

if we pass channel as <-chan in a func it means it's receive only and if we pass it as chan<- that means it's send only channel

```go
func receiverOnlyChan(res <-chan int) {
    for range res {
        fmt.Println("receive only cannot send to res")
    }
    // res <-5 not possible in this func scope
}

func sendOnlyChan(sen chan<- int) {
    for i := range 10 {
        sen <-i
    }
    // i <-res not possible in this func scope
}
```
