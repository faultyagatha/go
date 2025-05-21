package main

import (
	"fmt"
	"unsafe"
)

func printBytes(n int) {
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
}

func main() {
	s := make([]int32, 1000)
	var i32 int32
	var i16 int32
	var f32 float32
	var f64 float64
	var r rune
	var str string
	var p *int

	fmt.Println("Size of int16:", unsafe.Sizeof(i16))
	fmt.Println("Size of int32:", unsafe.Sizeof(i32))
	fmt.Println("Size of float32:", unsafe.Sizeof(f32))
	fmt.Println("Size of float64:", unsafe.Sizeof(f64))
	fmt.Println("Size of rune:", unsafe.Sizeof(r))
	fmt.Println("Size of str:", unsafe.Sizeof(str))
	fmt.Println("Size of p:", unsafe.Sizeof(p))

	fmt.Println("Size of []int32:", unsafe.Sizeof(s))
	fmt.Println("Size of [1000]int32:", unsafe.Sizeof([1000]int32{}))
	fmt.Println("Real size of s:", unsafe.Sizeof(s)+unsafe.Sizeof([1000]int32{}))
}
