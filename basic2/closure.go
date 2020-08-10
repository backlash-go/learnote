package main

import "fmt"

func main() {
/*
闭包

一个外层函数中  有内层函数，该内层函数会操作外层函数的局部变量（外层函数的参数， 外层函数定义的变量）
并且该外层函数的返回值就是内层函数，这个内层函数和外层函数的局部变量， 统称为闭包结构

外层函数的局部变量不会随着外层函数的调用而销毁，而是内层函数继续使用
 */
i1 := increment()
fmt.Println(i1)
fmt.Println(i1())
i2 := i1
fmt.Println(i2()) //2
fmt.Println(i2())  //3
fmt.Println(i2())  //4

i3 := increment()
fmt.Println(i3())//1
fmt.Println(i3())//2
fmt.Println(i3())//3

}


func increment() func()int{
	i:=0
	fun := func()int {
		i++
		return i
	}
	return fun
}
