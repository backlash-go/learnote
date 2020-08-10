package main

import "fmt"

func main() {
	/*
		pointer:  存储另一个变量的内存地址  的变量
	*/
//var a int = 10
	a := 10   //
	fmt.Printf("%T,%p\n", a, &a)

	//指针定义
	var apointer *int
	fmt.Println(apointer) //<nil>  空指针

	apointer = &a //apointer 指向了a的内存地址 0xc000098008
	fmt.Println(apointer)
	fmt.Printf("%p\n", &apointer) //apointer 自己的内存地址 0xc000092020
	fmt.Println(*apointer)        //获取a内存地址中存储的值 10
	a = 100                       //a变量的内存地址不变，
	fmt.Printf("%T,%p\n", a, &a)

	*apointer = 200 //通过指针修改变量的值
	fmt.Println(a)

	//指针的指针
	var p1 **int
	fmt.Println(p1)
	p1 = &apointer
	fmt.Println(p1)
	fmt.Println(**p1) //获取a的值

}
