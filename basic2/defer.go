package main

import (
	"fmt"
)

func main() {
	/*
		defer 词义 延迟 推迟
		在go语言中defer 关键字用来延迟一个函数或者方法的执行
		defer 的用法：
			对象的close，临时文件的删除
				文件.open()
				defer close（）

			go语言中关于异常的处理 使用panic（）和recover（）
			Panic 用于程序中断
			recover恢复程序执行， recover语法上要求必须在defer中执行

	defer fun1("hello") //多个defer  谁先被延迟谁最后    相当与栈 先进后出 LAST IN FIRST OUT
	fmt.Println("12345")
	defer fun1("world") //world 输出在程序的最后面
	fmt.Println("王二狗")

	// defer 函数的参数传递   函数调用时，参数就已经传递， 只是代码没有执行而已
	*/
	//a := 2
	//fmt.Println(a)
	//defer fun2(a)
	//a++
	//fmt.Println(a)
	//fmt.Println(fun3())

	fun3()

}
func fun3() int  {   // 有RETURN时 所有延迟函数执行完才会 RETURN返回值
	fmt.Println("hanshu")
	defer fun1("haha")
	return 0

}

func fun1(s string) {
	fmt.Println(s)


}

func fun2(n int) {
	fmt.Println(n)

}

