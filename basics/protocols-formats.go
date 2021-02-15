package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAddress() {
	resp, err := http.Get("www.faultyagatha.io")
	fmt.Println(resp, err)
}

//net.Dial("tcp", "faultyagatha.io:80")

type Person struct {
	name    string
	address string
	phone   string
}

func jsonMarsh() {
	person := Person{name: "joe", address: "a st.", phone: "911"}

	barr, err := json.Marshal(person) //returns JSON representation as []byte
	fmt.Println(barr, err)            //[123 125] <nil>

	var person2 Person
	err2 := json.Unmarshal(barr, &person2) //must fit JSON []byte
	fmt.Println(person2, err2)
}

func main() {
	getAddress()
	jsonMarsh()
}
