package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the number of intervals
	numberOfIntervals := 60

	// Create a slice of time intervals, each representing 1 minute
	timeIntervals := make([]int, numberOfIntervals)
	for i := range timeIntervals {
		timeIntervals[i] = i + 1
	}

	// Loop through the intervals
	for _, interval := range timeIntervals {
		fmt.Printf("Waiting for %d minute(s)...\n", interval)

		// Start a timer
		timer := time.Now()

		// Wait for the specified interval
		for time.Since(timer).Minutes() < float64(interval) {
			// Do nothing, just wait
		}

		// Stop the timer
		elapsed := time.Since(timer)
		fmt.Printf("Action performed after %d minute(s). Elapsed time: %v\n", interval, elapsed)
	}
}
