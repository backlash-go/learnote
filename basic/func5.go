package main

import "fmt"

func main() {
	/*
		作用域：变量可以使用的范围
			局部变量   函数内部定义的变量
			全局变量	  函数外部定义的变量
	*/
	n := 10 //定义在main函数中，所有n的作用域是main函数范围内
	fmt.Println(n)
	if a := 1; a <= 10 { //a的作用域只在分支语句
		fmt.Println(a)
		fmt.Println(n)
	}

	if b := 1; b <= 10 {
		n := 20
		fmt.Println(b)
		fmt.Println(n)
	}
	//fmt.Println(a) //undefined: a  //不能a，出了作用域
}

func fu1() {
num1 := 100
fmt.Println(num1)

}
