package main

import "fmt"

func main() {
	a := 99
	switch a {   //case 数值类型必须与SWITCH数值类型一致  ，case 是无序的数值是唯一的
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("input is false")
	}

	fmt.Println("main is over")
}