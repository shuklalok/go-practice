package main

import (
	"fmt"
	"time"
)

func main() {
	go count("apple")
	go count("banana")
}

func count(fruit string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, fruit)
		time.Sleep(time.Millisecond * 500)
	}
}

// OUTPUT

// ❯ go run 03-main-exits-early/main-exits-early.go
