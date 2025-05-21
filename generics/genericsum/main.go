package main

import "fmt"

// sumIntsOrFloats demonstrates a simple example of
// how to use generics in Go
func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// number is a type constraint
type number interface {
	int64 | float64
}

// sumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func sumNumbers[K comparable, V number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// You can omit type arguments in calling code when the Go compiler can infer
	// the types you want to use. The compiler infers type arguments from the types
	// of function arguments.
	fmt.Printf("Generic Sums: %v and %v\n",
		sumIntsOrFloats[string, int64](ints),
		sumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		sumNumbers(ints),
		sumNumbers(floats))
}
