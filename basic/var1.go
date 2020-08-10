package main

import "fmt"

func main()  {
	/*
	变量：variable
	概念：就是一小块内存用来存储数据
	使用：
		step1 变量的声明，也叫定义
		step2 变量的访问，赋值和取值
	*/
	//第一种 定义变量，然后赋值
	var num1  int
	num1 = 30
	fmt.Println(num1)
	fmt.Printf("num1的数值是: %d\n",num1)
	//写在一行
	var num2 int = 15
	fmt.Printf("num2的数值是: %d\n",num2)
	//第二种，类型推断
	var name = "lyq"
	fmt.Printf("name的数据类型是 %T, 数值是%s\n", name, name)
	//第三种，简短定义， 也叫简短声明
	sum := 100
	fmt.Printf("sum的数值是%s\n", sum)
	fmt.Println(sum)

	//多个变量同时定义
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	var m, n int
	m = 4
	n = 5
	fmt.Println(m, n)

	var n1, f1, s1 = 6, 3.14, "go"
	fmt.Println(n1, f1, s1)
	var (
		studentName = "席贤斌"
		age = 10
		sex = "男"
	)
	fmt.Printf("学生姓名: %s, 学生年龄: %d, 学生性别: %s", studentName, age, sex)
}
