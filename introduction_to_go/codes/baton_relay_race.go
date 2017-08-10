package main

import (
	"fmt"
	"time"
)

// start main OMIT
func main() {
	// Create an unbuffered channel // HL
	baton := make(chan int)

	// First runner to his mark // HL
	go Runner(baton)

	// Start the race // HL
	baton <- 1

	// Give the runners time to race // HL
	time.Sleep(500 * time.Millisecond)
}

// end main OMIT

// start Runner OMIT
func Runner(baton chan int) {
	var newRunner int
	// Wait to receive the baton // HL
	runner := <-baton
	// Start running around the track // HL
	fmt.Printf("Runner %d Running With Baton\n", runner)
	// New runner to the line // HL
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	// Running around the track // HL
	time.Sleep(100 * time.Millisecond)
	// Is the race over // HL
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		return
	}
	// Exchange the baton for the next runner // HL
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}

// end Runner OMIT

