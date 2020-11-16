package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Start working...")
	time.Sleep(3 * time.Second)
	fmt.Println("Work is done!")
	done <- true // Sends the value 'true' to the channel
}

func send(msgChan chan<- string, msg string) {
	msgChan <- msg // Sends string message to the channel
}

func receive(msgChan <-chan string) {
	fmt.Println(<-msgChan) // Prints the message received on channel
}

func main() {
	done := make(chan bool, 1) // Creates a channel of the type 'bool' with a capacity of 1 and that is an unbuffered channel, the channel can only send or receive a boolean value

	// Curiosity: A buffered channel has capacity of more than 1
	// channel := make(chan int, 5)
	// The channel can contain 5 elements before it is blocked temporarily

	go worker(done) // Start Go Routine
	<-done          // Discard the channel value without using it

	msgChan := make(chan string, 1) // Creates a channel of the type 'string' with a capacity of 1
	send(msgChan, "Hello, world!")
	receive(msgChan)
}
