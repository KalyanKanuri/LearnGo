# Channels

Channel is a Data Structure which is provided go, channels is a pipe for the data to flow from one goroutine to another

## UnBuffered Channels

UnBuffered channels are channels which cannot store values pushed by producer, the values pushed by the producer should immediately be consumed by consuemt or else the producer will wait forever until a consumer picks up the value and it results in deadlock

```go
// Making an Unbuffered Channel

ch := make(chan int)
```

## Buffered Channels

Unlike Unbuffered channels, Buffered channels has the capacity to store the values to a certain limit pushed by the producer when the storage in the channel finishes it behaves same as the unbuffered channels with this even though the consumer is not receiving values the producer can continue without moving to deadlock or without waiting forever for the consumer

```go
// Making a Buffered channel
ch := make(chan int, 3)

// this channel can hold 3 integer values even though the consumer is not pulling the values from the channel
```
