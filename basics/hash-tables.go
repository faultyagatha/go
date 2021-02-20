package main

import (
	"fmt"
	"strings"
)

//Vertex struct to demonstrate map example
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

//the most common way to create a map using make()
func mapMake() {
	var idMap map[string]int //key-val
	idMap = make(map[string]int)
	idMap["first"] = 1
	idMap["second"] = 2
	fmt.Println(idMap)
}

func mapLiteral() {
	idMap := map[string]int{"joe": 123}
	fmt.Println(idMap["joe"])
	idMap["jane"] = 456
	fmt.Println(idMap)
	delete(idMap, "joe")
	fmt.Println(idMap)
}

func mapIter() {
	idMap := map[string]int{"joe": 123, "jane": 456}
	for key, val := range idMap {
		fmt.Println(key, val)
	}
}

func mutateMap() {
	idMap := map[string]int{"joe": 123, "jane": 456}
	idMap["joe"] = 789 //mulated "joe"
	//add camila, abe
	idMap["camila"] = 101112
	idMap["abe"] = 1000
	fmt.Println(idMap)
	delete(idMap, "abe")
	fmt.Println(idMap)
	//test if a key is present
	elem, ok := idMap["camila"]
	fmt.Println("The value:", elem, "Present?", ok)
}

//returns a map of the counts of each “word” in the string
func wordCount(s string) map[string]int {
	//https://golang.org/pkg/strings/#Fields
	strArr := strings.Fields(s)
	newMap := make(map[string]int)
	for _, v := range strArr {
		newMap[string(v)]++
	}
	fmt.Println(newMap)
	return newMap
}

func main() {
	mapMake()
	mapLiteral()
	mapIter()
	mutateMap()
	wordCount("What a beautiful day")
	//how to use map with custom data type
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
