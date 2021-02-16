package main

import (
	"fmt"
	"reflect"
)

func arrays() {
	//initialised to 0 values
	var arr [5]int
	arr[0] = 2
	var arrLiteral [5]int = [5]int{1, 2, 3, 3, 4}
	fmt.Println(len(arrLiteral))
	var s []int
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

//array itteration range
func arrInteration() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	//two values are returned for each iteration.
	//the first is the index, and the second is a copy of the element at that index
	for i, v := range x {
		fmt.Printf("ind %d, val %d", i, v)
	}
}

//you can skip the index or value by assigning to _.
func skipped() {
	var y [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, _ := range y {
		fmt.Printf("ind %d, val %d", i)
	}
}

func main() {
	arrays()
	arrInteration()
	skipped()
}
