package main

import "fmt"

func main() {
	/*
		匿名函数： 就是没有名字的函数
		作用 ： go语言支持函数式编程， 将匿名函数作为另外一个函数的参数， 将匿名函数作为另外一个函数的返回值，形成闭包结构
	*/
	n1()
	n2 := n1
	n2()
	//匿名函数    定义匿名函数 直接调用  通常只能只用一次   赋值给变量就可以调用多次
	func () {
		fmt.Println("我是匿名函数")
	}()
	n3 := func () {
		fmt.Println("我是匿名函数")
	}
	n3()
	n3()
	//定义带参返回值的匿名函数
	func (a,b int)(int,int){
		fmt.Println(a,b)
		return a,b
	}(1,2)
}

func n1() {
	fmt.Println("a")
}
