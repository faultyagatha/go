package main

import "fmt"

//Vertex struct to demonstrate map example
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func mapMake() {
	var idMap map[string]int //key-val
	idMap = make(map[string]int)
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

func main() {
	mapMake()
	mapLiteral()
	mapIter()
	//how to use map with custom data type
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
