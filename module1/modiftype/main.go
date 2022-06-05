package main

import (
	"fmt"
	"strconv"
)

//
//type NewInt int
//type IntAlias = int

func main() {
	// 字符串与其他类型的转换
	// str 转 int
	newStr1 := "1"
	intValue, _ := strconv.Atoi(newStr1)
	fmt.Printf("%T,%d\n", intValue, intValue) // int,1

	// int 转 str
	intValue2 := 1
	strValue := strconv.Itoa(intValue2)
	fmt.Printf("%T, %s\n", strValue, strValue)
}
