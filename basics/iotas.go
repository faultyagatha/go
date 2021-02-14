package main

type Grades int

//the same as enums: starts at 1 and increments each val
const (
	A Grades = iota //1
	B               //2
	C               //3
	D               //4
	F               //5
)
