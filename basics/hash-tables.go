package main

import "fmt"

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
}
