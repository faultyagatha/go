package main

import "fmt"

//data flow in the direction of the arrow

//must be created before using
func makeChannel() chan int {
	//can be buffered: ch := make(chan int, 100)
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

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := makeChannel() //can be simply c := make(chan int)
	//distribute summing between two goroutines
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("x: ", x, "y: ", y, "sum: ", x+y)
}
