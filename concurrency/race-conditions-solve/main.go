// This sample program demonstrates how to solve race
// conditions in our programs using atomic.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int64

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup
)

// incCounter increments the package level counter variable.
func incCounter() {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Safely add 1 to a Counter
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	// Only one goroutine can mutate the value of 
	// a counter at a time.
	go incCounter()
	go incCounter()

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}