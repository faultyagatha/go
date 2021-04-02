package main

import "fmt"

/*Use Cases:
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
}
