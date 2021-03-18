package main

import (
	"fmt"
	"strings"
)

/* slice is an array-like data type:
- can be size-flexible, up to a size of an array
- a 'window' on an underlying array
- every slice has 3 properties:
	1. pointer: start of the slice
	2. length: the number of elements in the slice: len()
	3. capacity: maximum number of elements: cap()

	a[low : high] //includes 1 and excludes last elem
	var a [10]int
	a[0:10]
	a[:10]
	a[0:]
	a[:]

	IMPORTANT:
	A slice does not store any data, it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding elements of its underlying array.
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
//helps to create dynamically-sized arrays
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

func namesSlice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func tictactoe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceMethods() {
	var arr [10]int
	var sl []int
	fmt.Println("slice: %d", sl)  //slice: %d []
	fmt.Println("array: %d", arr) //array: %d [0 0 0 0 0 0 0 0 0 0]
	sl2 := append(sl, 10)
	fmt.Println("slice2: %d", sl2) //slice: %d []
	fmt.Println("slice1: %d", sl)  //slice: %d []
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	sliceExample()
	sliceLiteral()
	sliceMake()
	namesSlice()
	tictactoe()
	sliceMethods()
	a := []int{1, 2, 3}
	reverse(a)
	fmt.Println(a)
}
