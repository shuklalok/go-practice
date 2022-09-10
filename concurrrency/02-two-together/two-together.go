package main

import (
	"fmt"
	"time"
)

func main() {
	go count("apple")
	count("banana")
}

func count(fruit string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, fruit)
		time.Sleep(time.Millisecond * 500)
	}
}

// OUTPUT

// ❯ go run 02-two-together/two-together.go
// 0 banana
// 0 apple
// 1 apple
// 1 banana
// 2 banana
// 2 apple
// 3 apple
// 3 banana
// 4 banana
// 4 apple
// 5 apple
// 5 banana
