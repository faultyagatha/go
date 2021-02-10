package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

var c, python, java bool //declare vars and give them bool type

var x, y int = 1, 2

func main() {
	fmt.Println("Hello,  世界!")
	fmt.Println(add(2000, 21))
	var i int //auto initialisation to 0
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	k := 3 //short assignment statement can be used in place of a var declaration with implicit type (only inside function!!)
	fmt.Println(x, y, c, python, java, k)
}
