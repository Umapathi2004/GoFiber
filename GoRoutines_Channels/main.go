package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	for i := 0; i < 3; i++ {
		// Simulate work
		time.Sleep(time.Duration(id) * time.Second)
		// Send a message to the channel
		ch <- fmt.Sprintf("Worker %d: Completed task %d", id, i+1)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start multiple goroutines
	go worker(1, ch1)
	go worker(2, ch2)

	// Use select to receive messages
	for i := 0; i < 6; i++ { // Total tasks = 2 workers * 3 tasks each
		select {
		case msg := <-ch1:
			fmt.Println("From Worker 1:", msg)
		case msg := <-ch2:
			fmt.Println("From Worker 2:", msg)
		}
	}

	fmt.Println("All tasks completed.")
}
