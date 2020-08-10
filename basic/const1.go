package main

import (
	"fmt"
)

func main() {
	/*
		常量：
		1，概念 程序执行过程中数值不能改变, 尽量用驼峰
		2，语法
			1，显式类型定义

			2，隐式类型定义
		3，常数
			固定的数值，100 abc
	*/
	fmt.Println(100)

	const PATH string = "http://www.baidu.com"
	const pi = 3.14
	fmt.Println(PATH, pi)

	const c1, c2, c3 = 1, 2, 3
	const (
		name = "lyq"
		age  = "100"
		sex  = "nv"
	)
	//一组常量中如果某个常量没有初始值，类型与默认值与上一行一致
	const (
		a = 1
		b = 2
		c
		f
		d string = "xxb"
	)
	fmt.Println(a, b, c)
	fmt.Printf("%T, %d\n", c, c)
	fmt.Printf("%T, %d\n", b, b)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %d\n", f, f)
}
