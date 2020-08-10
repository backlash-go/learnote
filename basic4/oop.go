package main

import (
	"fmt"
)

func main() {
	/*
	oop : Object-oriented programming
		 模拟继承   通过嵌套 匿名字段
	模拟聚合    通过结构体嵌套， 没有匿名字段
	 */
	s1 := student{person{name:"xxb",age:13},"beijingdaxue"}
	s2 := student{person{name:"lll",age:123},"wuyixueyuan"}
	fmt.Println(s1,s2)
	fmt.Println(s1.name,s1.person.name)//xxb xxb  	匿名字段嵌套的结构体里面的字段name age是对于student 就是提升字段 可以简写

}
//person class
type person struct {
	name string
	age int
}

//sub class

type student struct {
	person  //匿名字段   嵌套的结构体里面的字段是提升字段
	school string
}
