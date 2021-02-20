package main

import (
	"fmt"
	"time"
)

//SYNCHRONISATION is using global events whose execution
//is viewed by all threads, simultaneously
//synchronisation is used to restrict bad interleavings

func main() {
	go fmt.Printf("New routine")
	//this is bad cause we make assumptions about time (non-deterministic)
	time.Sleep(100 * time.Millisecond) //hack to allow "New routine" to be printed
	fmt.Printf("Main routine")         // without the hack above, prints only "Main routine"
}
