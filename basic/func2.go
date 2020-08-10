package main

import "fmt"

var sd string = "as"
func main() {
	/*
	参数传递
		1 值传递    传递数组的副本
	          值类型数据，默认都是值传递， 包括 基本数据类型，数组，struct（结构体）
		2 引用传递  传递内存地址      切片， map, channel
	*/
	var arr1 = [4]int{1, 2, 3, 4}
	fmt.Println(arr1, sd)
	fun1(arr1)

	s1 := []int{5,6,7,8}
	fmt.Println(s1)
	fun2(s1)
	fmt.Println(s1)


}

func fun1(arr2 [4]int) {
	arr2[0] = 100
	fmt.Println(arr2)

}

func fun2(s2 []int)  {
	s2[0] = 100
	fmt.Println(s2)

}
