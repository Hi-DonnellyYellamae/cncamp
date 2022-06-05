package main

import (
	"fmt"
	"strings"
)

func main() {
	//var str string = "This is an example of a string"
	//fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	//fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	//var str1 string = "hello"
	//var str2 string = "hello,码神之路"
	//// 遍历
	//for i := 0; i < len(str1); i++ {
	//	fmt.Printf("ascii: %c %d\n", str1[i], str1[i])
	//}
	//for _, s := range str1 {
	//	fmt.Printf("unicode: %c %d\n ", s, s)
	//}
	//// 中文只能用 for range
	//for _, s := range str2 {
	//	fmt.Printf("unicode: %c %d\n ", s, s)
	//}
	tracer := "码神来了,码神bye bye"

	// 正向搜索字符串
	comma := strings.Index(tracer, ",")
	fmt.Println(",所在的位置:", comma)
	fmt.Println(tracer[comma+1:]) // 码神bye bye

	add := strings.Index(tracer, "+")
	fmt.Println("+所在的位置:", add) // +所在的位置: -1

	pos := strings.Index(tracer[comma:], "码神")
	fmt.Println("码神，所在的位置", pos) // 码神，所在的位置 1

	fmt.Println(comma, pos, tracer[5+pos:])
}
