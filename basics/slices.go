package main

import "fmt"

/* slice is an array-like data type:
- can be size-flexible, up to a size of an array
- a 'window' on an underlying array
- every slice has 3 properties:
	1. pointer: start of the slice
	2. length: the number of elements in the slice: len()
	3. capacity: maximum number of elements: cap()
*/

func sliceExample() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3] //b c
	s2 := arr[2:5] //c d e
	fmt.Println("s1: ", s1, "s2: ", s2)
	fmt.Println("s1 len: ", len(s1), "s2: ", len(s2))
}

func sliceLiteral() {
	sli := []int{1, 2, 3}
	fmt.Println("sli: ", sli)
	fmt.Println("sli cap: ", cap(sli))
}

//creates a slice with init vals of 0
func sliceMake() {
	sli := make([]int, 10) //init to 0 ith length 10
	fmt.Println("sli: ", sli)
	fmt.Println("sli cap: ", cap(sli))
	sli2 := make([]int, 10, 15) //init to 0 with length 10 and capacity up to 15
	fmt.Println("sli2 cap: ", cap(sli2))
}

//appends to a slice - can excede the size of an array - the memory will
//be reallicated as in C++ vectors
func sliceAppend() {
	sli := make([]int, 0, 3)
	sli = append(sli, 100)
}

func main() {
	sliceExample()
	sliceLiteral()
	sliceMake()
}
