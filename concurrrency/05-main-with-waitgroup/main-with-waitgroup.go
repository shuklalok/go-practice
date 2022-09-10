package main

import (
	"fmt"
	"sync"
)

func main() {
	// create a wait group, it is just a number.
	var wg sync.WaitGroup
	// add 1 goroutine to be waited for.
	wg.Add(1)
	// A wrapper function to help access the wg var.
	go func() {
		count("apple")
		// mark it done.
		wg.Done()
	}()
	// wait for the goroutine.
	wg.Wait()
}

func count(fruit string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, fruit)
	}
}

// OUTPUT

// ❯ go run 04-let-main-wait/let-main-main.go
// 1 apple
// 2 apple
// 3 apple
// 4 apple
// 5 apple
// 6 apple
