package main

import (
	"errors"
	"fmt"
	"os"
)

type error interface {
	Error() string
}

func readFile() {
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}

//use it like this:

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		//this is a common way to return double values
		//with one being an error and another one is empty
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

type SyntaxError struct {
	msg    string // description of error
	Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }

type appError struct {
	Error   error
	Message string
	Code    int
}
