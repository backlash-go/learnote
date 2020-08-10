package main

import "fmt"

func main() {
	/*
	流程控制结构：
		1 顺序结构
		2 选择结构
		3 循环结构
	 */

	//if语句
	// 1
	var a int = 11
	if a > 10 {
       fmt.Println("yes")
	}
	//2
	if a < 10 {
		fmt.Println("a 小于10")
	}else {
		fmt.Println("a 大于10")

	}
	//3
	if a < 10 {
		fmt.Println("a 小于10")
	}else if a > 10 {
		fmt.Println("a over 10")

	}else {
		fmt.Println("nothing to know")

	}

		fmt.Println("main is over")
}
