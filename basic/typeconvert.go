package main

import (
	"fmt"
)

func main() {
	//数据类型转换需要兼容类型， 转换针对变量需要手动， 常量自动转型
	var a int8
	a = 10
	var b int16
	b = int16(a)
	fmt.Println(a, b)

	f1 := 4.83
	var c  int
	c = int(f1)
	fmt.Println(c, f1) //4 4.83

	sum1 := f1 + 100
	fmt.Printf("%T,%d\n",sum1,sum1)
}
