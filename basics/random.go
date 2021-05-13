package main

import (
	"math/rand"
	"strconv"
)

//generates a string with random numbers
func randomNumbersStr(n int) string {
	nums := ""
	for i := 0; i < n; i++ {
		r := rand.Intn(5000)
		nums = nums + strconv.Itoa(r)
	}
	return nums
}

//generates an array of random numbers
func randomNumbers(n int) []int {
	nums := []int{}
	for i := 0; i < n; i++ {
		r := rand.Intn(5000)
		nums = append(nums, r)
	}
	return nums
}

//a seed string for random letters
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

//generates an array of random letters
func randomLetters(n int) []rune {
	l := make([]rune, n)
	for i := range l {
		l[i] = letters[rand.Intn(len(letters))]
	}
	return l
}
