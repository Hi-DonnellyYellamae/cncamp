package main

import "fmt"

func main() {
	//myArray := [5]int{1, 2, 3, 4, 5}
	//myslice := myArray[1:3]
	//fmt.Printf("mySlice %+v\n", myslice)
	//fullslice := myArray[:]
	//remove3rdItem := deleteItem(fullslice, 2)
	//fmt.Printf("remove3rdItem %+v\n", remove3rdItem)
	s1 := "localhost:8080"
	fmt.Println(s1)
	strByte := []byte(s1)
	strByte[len(s1)-1] = '1'
	fmt.Println(strByte)
	s2 := string(strByte)
	fmt.Println(s2)
}

//func deleteItem(slice []int, index int) []int {
//	return append(slice[:index], slice[index+1:])
//}
