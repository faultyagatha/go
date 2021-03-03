package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

//the string type in Go is an array/slice of bytes
//(usually UTF-8 strings but may also be just binary data)

//the name “rune” is another name for a Unicode code point
//and is a 32 bit integer

//Go strings are immutable and behave like read-only byte slices
//(with a few extra properties).
//To update the data, use a rune slice instead (see example in main).

//******CASTING TO RUNES
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

//better option, as of 2020
func sortStringBySlice(s string) {
	s1 := "eidbaooo"
	runeSlice := []rune(s1)
	fmt.Println(string(runeSlice))
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})

	fmt.Println(string(runeSlice))
}

//check if the string has only letters
func hasLetters(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

func main() {
	//***** IMMUTABLE
	buf := []rune("hello")
	buf[0] = 'H'
	st := string(buf)
	fmt.Println(st) // "Hello"

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
