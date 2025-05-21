package main

import (
	"fmt"
	"math"
)

//FUNCTIONS IS GO ARE FIRST-CLASS:
//treated like other types (think Javascript)
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func withReturn(x, y float32) float32 {
	return x + y
}

func doubleReturn(x int) (int, int) {
	return x, x + 1
}

//closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
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

//2. not the best way to do in GO but possible
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

//function defines a function and returns a function
func makeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

//VARIADIC FUNCTION
//takes an arbitrary num of ints or
//we can also pass a slice of ints
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

//we can declare a method on Go types
//we can only declare a method with a receiver whose type is defined in
//the same package as the method.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// recursive function to calculate gcd
// using Euclidian algorithm
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//DEFERRED FUNCTION
//typically for clean up
//called only when the surrounding function completes
//IMPORTANT: the arguments are evaluated immediately
//but the call of the function is deferred

func main() {
	defer fmt.Println("Bye!")
	funcVar = incFn
	fmt.Print(funcVar(1))

	y := withReturn(19.9, 1.1)
	fmt.Println(y)

	a, b := doubleReturn(3)
	fmt.Println(a, b)

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

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

	dist1 := makeDistOrigin(0, 2)
	dist2 := makeDistOrigin(2, 3)
	fmt.Println(dist1(2, 2)) //2
	fmt.Println(dist2(2, 2)) //1

	fmt.Println(getMax(1, 3, 6, 4)) //6
	vslice := []int{1, 3, 6, 4}
	fmt.Println(getMax(vslice...)) //6: ...suffix is necessary

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	res:= gcd(874, 1944)
	fmt.Printf("gcd %d\n", res)
}
