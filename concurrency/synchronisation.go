package main

import (
	"fmt"
	"sync"
)

//SYNCHRONISATION is using global events whose execution
//is viewed by all threads, simultaneously
//synchronisation is used to restrict bad interleavings

func foo(wg *sync.WaitGroup) {
	fmt.Printf("New routine")
	wg.Done() //must be called when routine is completed
}

func main() {
	//sync package can synchronise between goroutines
	var wg sync.WaitGroup
	wg.Add(1) //instance of sync contains internal counter
	go foo(&wg)
	wg.Wait()                  //wait for Done() from foo
	fmt.Printf("Main routine") //main will not be executed until the counter is complete
}
