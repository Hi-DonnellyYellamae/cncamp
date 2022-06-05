package main

import "fmt"

func main() {
	var mapList map[string]int
	var mapAssigned map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	mapAssigned = mapList
	mapAssigned["two"] = 3
	fmt.Println(mapList["one"])
	fmt.Println(mapList["two"])
	fmt.Println(mapList["ten"])

}
