package main

import "fmt"

func main() {
	/*
	1 算术运算符   + - * / % ++ --
	2 关系运算符   == != > <  >=  <=
	3 逻辑运算符
	 */
	a := 10
	b := 3
	div := a / b
	mod := a % b
	fmt.Printf("%d / %d = %d\n", a,b,div)
	fmt.Printf("%d %% %d = %d\n", a,b,mod)

	c := 3
	//c++
	c--
	fmt.Println(c)
}
