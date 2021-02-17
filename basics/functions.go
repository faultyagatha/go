package main

import "fmt"

func withReturn(x, y float32) float32 {
	return x + y
}

func doubleReturn(x int) (int, int) {
	return x, x + 1
}

//GO PASSES ARGUMENTS BY VALUE
//helps data incapsulation
//but for large objects, it takes time to copy
func passByValue(y int) {
	y = y + 1
}

//to pass large objects or
//to have a possibility to alter the argument
//use a pointer
//but there is no data encapsulation
func passByReference(y *int) {
	*y = *y + 1
}

//1.
func passArraysByValue(x [3]int) int {
	return x[0]
}

//2. not the best way to do in GO
func passArraysByRef(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

//3. the best way in GO is to use slice
//slice stores 3 things:
//a pointer to the start of the arr, length, capacity
//when we pass a slice, it's still by value
//but a function gets a pointer to the array
func passArraysBySlice(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	y := withReturn(19.9, 1.1)
	fmt.Println(y)

	a, b := doubleReturn(3)
	fmt.Println(a, b)

	x := 2
	passByValue(x)
	fmt.Println(x) //still 2

	passByReference(&x)
	fmt.Println(x) // 3

	ar1 := [3]int{1, 2, 3}
	fmt.Println(passArraysByValue((ar1))) // 1

	ar2 := [3]int{1, 2, 3}
	passArraysByRef(&ar2)
	fmt.Println(ar2)

	ar3 := []int{1, 2, 3} //this is a slice
	passArraysBySlice(ar3)
	fmt.Println(ar3)
}
