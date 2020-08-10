package main

import "fmt"

func main() {
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	/*
		iota 是一个特殊的常量，可以被编译器自动修改的常量
			每当定义一个const,iota的初始值为0
			每当定义一个常量，就会自动累加1
			直到下一个const出现，清零
		简写
		const (
		  	a = iota
			b
			c
		)
	*/
	const (
		d = iota
		e = iota
	)

	const (
		male = iota //0
		female     //female = iota 1
		unkonw     //unkonw = iota 2
	)
	fmt.Println(a, b, c)
	fmt.Println(d, e)
	fmt.Println(male, female, unkonw)
}
