package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printAlphabet(isLower bool, wg *sync.WaitGroup) {
	fmt.Println("\nStart Goroutines.\n")
	// Schedule the call to Done to tell the main we are done
	defer wg.Done()

	if isLower {
		for i := 0; i < 10000; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	} else {
		fmt.Println("\nStart Shorter Routine.\n")
		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println("\nEnd Shorter Routine.\n")
	}
}

func main() {
	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// main need to wait for the goroutines to finish
	var wg sync.WaitGroup
	// add 2 goroutines to the wg
	wg.Add(2)

	// these two goroutines will run concurrently
	// and will be swapped from time to time
	// (see the output)
	go printAlphabet(false, &wg)
	go printAlphabet(true, &wg)

	fmt.Println("\nWaiting to Finish.")
	wg.Wait()
	
	fmt.Println("\nTerminating main.")
}
