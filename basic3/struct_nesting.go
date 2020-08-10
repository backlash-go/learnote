package main

import (
	"fmt"
)

func main() {
	/*
	结构体嵌套   一个结构体的字段是一个结构体
	 */

	b1 := book{bookname:"English",price:4.23}
	s1 := st{name:"xxxb",age:18,bo:b1}
	fmt.Println(s1)
	fmt.Println(s1.bo.bookname)
	fmt.Printf("%T,%T\n",s1,b1)

	s2 := st{name:"lyq",age:19,bo:book{"go",90}}
	s3 := st{"sb",129,book{"gso",920}}
	fmt.Println(s2,s3)

	b4 := book{"python",76.9}
	s4 := st2{"a",19,&b4}
	fmt.Println(b4)
	fmt.Println(s4,*s4.bo)
	s4.bo.bookname="22"
	fmt.Println(b4,*s4.bo)



}

type book struct {
	bookname string
	price float64
}

type st struct {
	name string
	age int
	bo book
}

type st2 struct {
	name string
	age int
	bo *book //book 内存地址  会发生引用传递
}
