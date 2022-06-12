package main

import "fmt"

func main() {
	array := [...]string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("%v\n", array)
	for i := 0; i < len(array); i++ {
		switch i {
		case 2:
			array[i] = "smart"
		case 4:
			array[i] = "strong"
		default:
			continue
		}
	}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

}
