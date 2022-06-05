package main

import "fmt"

func main() {
	//var arr = [3]int{1, 2, 3}
	//for index, value := range arr {
	//	fmt.Printf("索引：%d，值：%d \n", index, value)
	//}
	//var strList []string
	//strList = append(strList, "lsddx")
	//fmt.Println(strList)
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{3, 4, 5}
	//copy(slice2, slice1)
	//fmt.Println(slice2)
	//fmt.Println(slice1)
	//copy(slice1, slice2)
	//fmt.Println(slice2)
	//fmt.Println(slice1)
	const elementCount = 1000
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	refData := srcData
	copyData := make([]int, elementCount)
	copy(copyData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[elementCount-1])
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
