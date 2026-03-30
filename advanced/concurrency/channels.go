package concurrency

import (
	"fmt"
)

func sendMsg(ch chan string, chanType string) {
	msg := fmt.Sprintf("Hello from sendMsg() using channel of channel type %s", chanType)
	ch <- msg // sending data using channel
}

func ChannelsInGo() {
	fmt.Println("-- Channels In Go --")
	ch := make(chan string)     // unbuffered channel
	ch2 := make(chan string, 3) // Buffered channel with 3 blocks of memory
	go sendMsg(ch, "unbuffered")
	go sendMsg(ch2, "buffered")

	select {
	case msg := <-ch: // received msg from channel
		fmt.Println(msg)
		close(ch) // closing channel after work done
	case msg2 := <-ch2:
		fmt.Println(msg2)
		close(ch2)
	}
	// msg = <-ch
	// if we try to receive another msg from the channel it will not give the same msg
	// the current goroutine will wait forever for the other goroutine to send the msg
	// but the goroutine is already completed it's execution so it results in deadlock
}
