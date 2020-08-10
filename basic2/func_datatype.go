package main

import "fmt"

func main() {
	/*
		基本数据类型，  int float bool string
		复合数据类型    array slice map function pointer  struct interface...
		函数的数据类型   func(参数列表的类型)(返回值列表的类型)

		函数的本质：是复合类型， 一种特殊的变量


	*/
	fmt.Printf("%T,%T,%T\n", q1, q2, q3) //func(),func(int) int,func(int, int, float64) (int, int)
	fmt.Println(q1)                      //0x109b210

	//定义一个函数类型的变量
	var c func(int) int //0x109b260
	fmt.Println(c)      //0x109b260  <nil>
	c = q2              //引用赋值 变量c指向q2的内存地址
	fmt.Println(c(1))   //0  调用c函数与q2函数返回结果一样

	//
	re1 := q2           //函数变量的赋值 引用赋值
	re2 := q2(2)        // 把函数的返回值传递给re2
	fmt.Println(re1, re2)
}

func q1() { //开辟一个内存  函数代码存在内存  q1就是一个特殊的变量指向内存地址0x109b210  加（）访问执行q1变量地址存储的代码

}

func q2(n int) int {
	return 0

}

func q3(a, b int, c float64) (int, int) {
	return 0, 0

}
