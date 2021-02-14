package main

import "fmt"

var trueFalse bool

var hello string

//int int8  int16  int32  int64
var simpleInt int

// uint uint8 uint16 uint32 uint64 uintptr
var unsignedInt uint

var uint8Alias byte // alias for uint8

var int32Alias rune // alias for int32
// represents a Unicode code point

var smalFloat float32  //6 digits of precision
var largeFloat float64 //15 digits of precision

var complexNum complex64
var veryLargeCmplexNum complex128

//********* INIT VALUES
/** uninitialised variables will have 0 values:
- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings
*/

//********* CASTING T(v)
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// i := 42
// f := float64(i)
// u := uint(f)

//********* IMPORTANT: assignment between items of different type requires an explicit conversion!!

//********* TYPE INFERENCE:
var x int

//y := x // y is an int, only possible inside of a function

const World = "世界" //cannot be declared using := syntax

//********* SCOPING:
var xScoped int = 4

func f1() {
	fmt.Printf("%d", xScoped)
}

func f2() {
	var xScoped int = 5
	fmt.Printf("%d", xScoped)
}

//********* POINTERS:
func main() {
	var x1 int = 1
	var y1 int
	var ip *int //declared as a pointer to an int
	ip = &x1
	y1 = *ip
	fmt.Println("x1 ", x1)
	fmt.Println("y1 ", y1)

	//new() function creates a var and returns a pointer to the var
	var x2 *int = new(int) //or x2 := new(int)
	fmt.Println("x2 ", x2) //prints some address: 0xc000012080
	*x2 = 3
	fmt.Println("x2 ", *x2) //prints 3
	f1()
	f2()
}
