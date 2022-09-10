package main

import (
	"fmt"
	"time"
)

func main() {
	count("apple")
	count("Banana")
}

func count(fruit string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, fruit)
		time.Sleep(time.Millisecond * 500)
	}
}

// OUTPUT:

// ❯ go run simple_goroutine/methods-one-by-one.go
// 0 apple
// 1 apple
// 2 apple
// 3 apple
// 4 apple
// 5 apple
// 0 Banana
// 1 Banana
// 2 Banana
// 3 Banana
// 4 Banana
// 5 Banana
