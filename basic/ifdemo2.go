package main

import "fmt"

func main() {
	if num := 3; num >= 4 {  //num 作用域是if结构内
		fmt.Println("num 是个正数")
	}else if num < 4 {
		fmt.Println("num is fushu")
	}
}
