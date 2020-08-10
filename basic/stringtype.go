package main

import (
	"fmt"
)

func main() {
	//1，定义字符串
	var s1 string
	s1 = "xixianbin"
	fmt.Println(s1)

	//2，区别  'A' "A"
	var v1 = 'A'
	var v2 = "A"
	fmt.Printf("%T, %d\n", v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	v3 := '中'
	fmt.Printf("%T,%d,%c,%q\n", v3, v3, v3, v3)

	//3，转义字符  \    \n    \t
	fmt.Println("\"xixianbin\"")
	fmt.Println(`"sd"`)
}
