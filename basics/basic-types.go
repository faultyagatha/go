package main

var trueFalse bool

var hello string

var simpleInt int

//int8  int16  int32  int64
var unsignedInt uint

//uint8 uint16 uint32 uint64 uintptr

var uint8Alias byte // alias for uint8

var int32Alias rune // alias for int32
// represents a Unicode code point

var smalFloat float32
var largeFloat float64

var complexNum complex64
var veryLargeCmplexNum complex128

//INIT VALUES
/** uninitialised variables will have 0 values:
- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings
*/

//CASTING T(v)
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// i := 42
// f := float64(i)
// u := uint(f)

//IMPORTANT: assignment between items of different type requires an explicit conversion!!

//TYPE INFERENCE:
var x int

//y := x // y is an int, only possible inside of a function

const World = "世界" //cannot be declared using := syntax
