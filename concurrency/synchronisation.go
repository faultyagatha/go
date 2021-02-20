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

//Sync.Once: handy for initialisation
var on sync.Once

func setup() {
	fmt.Printf(" Init")
}

func doStuff() {
	on.Do(setup)
	fmt.Printf(" Hello")
}

//synchronisation causes different goroutines depend on each other
//ch <- 1
//x := <- ch
//x depends on ch
//DEADLOCK is a circular dependency
func doChannelStuff(c1 chan int, c2 chan int, wg *sync.WaitGroup) {
	<-c1    //read to write onto first channel
	c2 <- 1 //wait to read from second channel
	wg.Done()
}

func main() {
	//sync package can synchronise between goroutines
	var wg sync.WaitGroup
	wg.Add(1) //instance of sync contains internal counter
	go foo(&wg)
	wg.Wait()                   //wait for Done() from foo
	fmt.Printf(" Main routine") //main will not be executed until the counter is complete
	doStuff()

	//how to create deadlock
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go doChannelStuff(ch1, ch2, &wg)
	go doChannelStuff(ch2, ch1, &wg) //here is a circular dependency
}
