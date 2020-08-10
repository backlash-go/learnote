package main

import "fmt"

var a int = 10
//b := 20   语法错误 syntax error: non-declaration statement outside function body
func main()  {
	var num int
	num = 100
	fmt.Printf("num的数值是%d,地址是%p\n", num, &num)
	num = 200
	fmt.Printf("num的数值是%d,地址是%p\n", num, &num)

    /*
    注意事项
    1，变量必须先定义才能使用
    2，go是静态语言，要求变量类型与赋值类型必须一致
    3，变量名不能冲突
    4，简短定义方式，左边的变量必须至少有一个是新的
    5，简短定义方式，不能定义全局变量
    6，变量的0值也叫默认值
    7，变量定义了就要使用， 否则无法通过编译
     */
    num, num1 := 200, 300  //no new variables on left side of :=
	fmt.Println(num, num1, a)
	//变量默认值
	/*
	整型，浮点型默认值为0
	字符，默认值为了""
	布尔，默认值为false
	 */
    var m int
    var n float64 //0.0 > 0
	var s  string //空字符串
	var bo  bool  //默认值为false
    fmt.Println(m, n, s, bo)
}
