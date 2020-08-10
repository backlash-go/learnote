package main

import "fmt"

func main() {
	/*
		高阶函数， 把一个函数作为另外一个函数的参数
		func1 func2  func1 作为func2的参数
		func2 就叫高阶函数
		func1 回调函数


		回调函数：

	*/
	fmt.Println(Operation(10, 20, subtraction))

	multiplication := func(a, b int) int { //使用匿名函数做乘法
		return a * b
	}
	fmt.Println(Operation(10, 20, multiplication))

	var division func(int, int) int
	division = func(a, b int) int {
		return a / b
	}
	fmt.Println(Operation(20, 5, division))

}

//加法运算

func add(a, b int) int {

	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func Operation(a, b int, fun func(int, int) int) int {
	r := fun(a, b)
	return r

}
