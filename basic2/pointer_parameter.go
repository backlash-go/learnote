package main

import (
	"fmt"
)

func main() {
	/*
	指针作为参数

	参数的传递， 值传递， 引用传递
	 */
	a:=10
	fmt.Println(a)
	po1(a)
	fmt.Println(a)

	po2(&a)
	fmt.Println(a)

	ar := [4]int{1,2,3,4}
	fmt.Println(ar)
	arra1(ar)
	fmt.Println(ar)

    arra2(&ar)
	fmt.Println(ar)

}

func arra1(a [4]int)  {
	a[0] = 100
	fmt.Println(a)


}

func arra2(a *[4]int)  {
	a[0] = 100
	fmt.Println(*a)

}

func po1(num int)  {
	num = 100
	fmt.Println(num)

}

func po2 (p1 *int){
	*p1 = 200
	fmt.Println(*p1)

}