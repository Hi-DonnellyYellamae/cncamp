package main

import "fmt"

func main() {
	const a = string("Hello")
	var b = string("world")
	fmt.Println(b)
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fullString := "Hello World!"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}
}
