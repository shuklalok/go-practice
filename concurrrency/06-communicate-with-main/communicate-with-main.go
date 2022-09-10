package main

import (
	"fmt"
	"time"
)

func main() {
	// Make a channel to communicate.
	c := make(chan string)
	// pass the channel.
	go count("apple", c)
	// --------------------------------------------------------
	// // for receiving all the messages,
	// // otherwise there will be only one received.
	// for {
	// 	// receiving is a blocking call
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	// --------------------------------------------------------
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(fruit string, c chan string) {
	for i := 0; i <= 5; i++ {
		// sending is a blocking call
		c <- fruit
		time.Sleep(time.Millisecond * 500)
	}
	// close the channel explicitly to avoid deadlock
	close(c)
	// sender should close, but never the receiver
}
