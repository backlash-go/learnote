package main

import "fmt"

func main() {
	/*
	数组指针，   一个指针， 数组的地址
	指针数组，
	 */

	array1 := [4]int{1,2,3,4}
	fmt.Println(array1)
	fmt.Printf("%p\n",&array1)

	//创建指针
	var p1 *[4]int
	p1 = &array1
	fmt.Printf("%p, %p\n",p1,&p1)  //&p1 指针自己的地址     p是数组的内存地址
	fmt.Println(*p1,p1) //[1 2 3 4]    &[1 2 3 4]数组的指针

	//根据指针修改数组
	(*p1)[0] = 100
	fmt.Println(array1)
	//简写方法
	p1[1] = 900
	fmt.Println(array1)


	//指针数组
	a:=1
	b:=2
	c:=3
	d:=4
	array2 := [4]int{a,b,c,d}
	array3 := [4]*int{&a,&b,&c,&d}
	fmt.Println(array2,array3)

	array2[0] =100
	fmt.Println(array2)
	fmt.Println(a)
	*array3[0] = 200
	fmt.Println(array3)
	fmt.Println(a)

}
