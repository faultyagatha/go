package main

import (
	"fmt"
	"math"
	"time"
)

func ifLoop() {
	var x int = 5
	if x > 5 {
		fmt.Println(x)
	}
}

//short if loop
func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		//variables declared by the statement are only in scope until the end of the if.
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//Go's while loop
func whileLoop() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

//no break in the end, it'll break automatically
//Go's switch cases need not be constants,
//and the values involved need not be integers
func switchFlow(x int) {
	switch x {
	case 1:
		fmt.Printf("case1")
	case -1:
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
	}
}

func multiSwitch(score int) {
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
}

//tagless switch
func taglessSwitch(x int) {
	switch {
	case x > 1:
		fmt.Printf("case1")
	case x < -1:
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
	}
}

/* Defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately,
but the function call is not executed until the surrounding function returns.
Deferred function calls are pushed onto a stack. When a function returns,
its deferred calls are executed in last-in-first-out order.
*/
func usingDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
