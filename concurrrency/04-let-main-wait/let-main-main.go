package main

import (
	"fmt"
	"time"
)

func main() {
	go count("apple")
	go count("banana")

	// main waits for 2 seconds
	time.Sleep(time.Second * 2)

	// main waits forever
	// for{}

	// main waits for user input
	// fmt.Scanln()
}

func count(fruit string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, fruit)
		time.Sleep(time.Millisecond * 500)
	}
}

// OUTPUT

// ❯ go run 004-let-main-wait/let-main-main.go
// 0 banana
// 0 apple
// 1 banana
// 1 apple
// 2 apple
// 2 banana
// 3 banana
// 3 apple
