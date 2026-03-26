package concurrency

import (
	"fmt"
)

func sendMsg(ch chan string) {
	msg := "Hello from sendMsg() using channel"
	ch <- msg // sending data using channel
}

func ChannelsInGo() {
	fmt.Println("-- Channels In Go --")
	ch := make(chan string)
	go sendMsg(ch)

	msg := <-ch // received msg from channel
	fmt.Println(msg)

	// msg = <-ch
	// if we try to receive another msg from the channel it will not give the same msg
	// the current goroutine will wait forever for the other goroutine to send the msg
	// but the goroutine is already completed it's execution so it results in deadlock
}
