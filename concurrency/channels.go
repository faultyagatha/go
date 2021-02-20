package main

import "fmt"

//channels is a way of communicating between coroutines
//data flow in the direction of the arrow -> or <-
//channel communication is synchronous
//blocking is the same as waiting for communication
//unbuffered channel cannot hold data in transit
//sender channel blocks until the data is received somewhere
//receiver channel blocks until the data is sent from sender

//must be created before using
func makeChannel() chan int {
	//unbuffered
	ch := make(chan int)
	return ch
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//buffering allows for independence from sender and receiver
//operating at exactly the same speed
func bufferedChannel() {
	//define capacity
	//we can do 100 sends and still not block
	//a buffer can hold the object
	//it'll start to block it's full: 100+1
	ch := make(chan int, 100)
	fmt.Println(ch)
}

//we can close a channel to indicate that no more values will be sent
//common in looping through channel to terminate a range loop
//closing is only necessary when the receiver must be told
//there are no more values coming
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		fmt.Println("x:", x, "y:", y, "sum:", x+y)
	}
	close(c)
}

//we can select the data to use (first-come first-served)
//we can use <-abort channel or quit channel to quit select
//default will be used if no other case is ready
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		//select blocks until one of its cases can run
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("I'm ready but other are not")
		}
	}
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := makeChannel() //can be simply c := make(chan int)
	//distribute summing between two goroutines
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("x: ", x, "y: ", y, "sum: ", x+y)

	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	//receives values from the channel repeatedly until it's closed
	for i := range ch {
		fmt.Println(i)
	}
}
