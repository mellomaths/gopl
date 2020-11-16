package main

import (
	"fmt"
	"sync"
	"time"
)

func worker() {
	fmt.Println("Worker finished!")
}

func wgWorker(wg *sync.WaitGroup) {
	fmt.Println("wgWorker finished!")
	wg.Done() // Decrease the counter by 1, indicate the termination of a go-routine
}

func main() {
	go worker() // Creates a Go Routine and calls worker()

	time.Sleep(1 * time.Second) // Achieves synchronisation so that main() waits to the end of worker() to complete its execution
	fmt.Println("Main finished!")

	var wg sync.WaitGroup // Forces a goroutine to wait for other goroutines
	wg.Add(1)             // Increase the Waitgroup counter
	go wgWorker(&wg)
	wg.Wait() // Blocks until the counter is 0.
}
