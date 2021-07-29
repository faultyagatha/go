package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadAll(readers []io.Reader) (string, error) {
	var sb strings.Builder
	for _, r := range readers {
		_, err := io.Copy(&sb, r)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

func ReadAll2(readers ...io.Reader) (string, error) {
	var sb strings.Builder
	for _, r := range readers {
		_, err := io.Copy(&sb, r)
		if err != nil {
			return "", err
		}
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

// ---------------
// Example of generics
// that could solve the problem.
// Generics is not yet in Go
//  ---------------
/*
func ReadAllGeneric[R io.Reader](readers []R) (string, error) {
  var sb strings.Builder
  for _, r := range readers {
    _, err := io.Copy(&sb, r)
    if err != nil {
      return "", err
    }
    sb.WriteString("\n")
  }
  return sb.String(), nil
}
*/

func main() {

	one := strings.NewReader("Hello, world!")
	two := strings.NewReader("How are you?")

	// ---------------
	// The code below will not work:
	// compilation error: cannot use wontWork (type []*strings.Reader)
	// as type []io.Reader in argument to ReadAll
	//  ---------------
	// wontWork := []*strings.Reader{one, two}
	// out, err := ReadAll(wontWork)

	// ---------------
	// The code below will work
	// because the type []io.Reader
	// is a concrete type in Go
	// and a type like []*strings.Reader
	// isn't the same in the compiler's eyes
	//  ---------------
	willWork1 := []io.Reader{one, two}
	out1, err1 := ReadAll(willWork1)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(out1)

	// ---------------
	// The code below will work
	// The only way to take a slice
	// of *strings.Reader and pass them
	// into a function that expects io.Readers
	// is to create a new slice and pass that in.
	//  ---------------
	wontWork := []*strings.Reader{one, two}
	var willWork2 []io.Reader
	for _, sb := range wontWork {
		willWork2 = append(willWork2, sb)
	}

	// This will work with either version of ReadAll
	out2, err2 := ReadAll2(willWork2...)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(out2)

	//example of using generics
	//worksNow := []*strings.Reader{one, two}
	//out, err := ReadAllGeneric(worksNow)
}
