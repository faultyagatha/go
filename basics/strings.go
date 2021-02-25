package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//***** MULTILINE
	str1 := `This is a
	multiline
	string.`
	fmt.Println(str1)

	//***** CONCAT
	// This works fine
	str := "abc"
	str = str + "def"

	// This is more efficient if you are combining lots of strings
	var sb strings.Builder
	sb.WriteString("abc")
	sb.WriteString("def")

	//***** CASTING TO NUMBER
	// This does NOT work as you may expect
	number := 123
	str2 := string(number)
	fmt.Println(str2) // Outputs: E or {

	// We need to use a function like strconv.Itoa or fmt.Sprintf
	strA := strconv.Itoa(number)
	fmt.Println(strA) // Outputs: 123

	strB := fmt.Sprintf("%d", number)
	fmt.Println(strB) // Outputs: 123

	//***** CASTING TO BYTE SLICES AND VICE-VERSA
	original := "this is a string"
	var b []byte
	b = []byte(original) // convert the string to a byte slice

	var s string
	s = string(b) // convert the byte slice to a string
	fmt.Println(s)
}
