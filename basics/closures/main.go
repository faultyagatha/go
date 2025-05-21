package main

import (
	"fmt"
	"time"
)

/* Use Cases:
- keep the local scope (as in the example below)
- often used with the standard libs
- passing middleware function
- deferring work
*/

func newCounter() func() int {
	n := 0 //scope of n
	return func() int {
		n++
		return n
	}
}

func main() {
	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// ----- 
	// Below is the unexpected behaviour until v1.21.
	// 1. The loop variable gets captured by func literal 
	// and produces unexpected result when used with a closure

	var prints []func()
	for i := 0; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println("I am a closure ", i) })
	}
	for _, print := range prints {
		// Prints 4 times: I am a closure  4
		print()
	}

	// 2. The loop variable i gets captured by func literal 
	// and produces unexpected result when used with goroutine

	for i := 0; i < 3; i++ {
		go func() {
			// Prints (in some arbitrary order) 3 times: 
			// I am goroutine 3
			fmt.Println("I am goroutine", i) 
		}()
	}

	// Fixes before v.1.21
	for i := 0; i < 3; i++ {
		// Copy the loop variable into a per-iteration variable of the same name
		i := i 
		go func() {
			// Prints (in some arbitrary order) something 
			// like: I am goroutine 0, 1, 2
			fmt.Println("I am goroutine", i) 
		}()
	}

	time.Sleep(100 * time.Millisecond)
}
