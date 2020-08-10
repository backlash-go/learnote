package main

import (
	"fmt"
)

func main() {

	/*
	匿名结构体 ： 就是没有名字的结构体
	匿名字段： 一个结构体的字段没有名字
	匿名函数：
	 */
	s1 := student{"lyq",18}
	fmt.Println(s1)
	fmt.Printf("%s\n",s1)

	//anonymous function
	func (){
		fmt.Println("hello")
	}()
	//anonymous struct
	s2 := struct {
		name string
		age int
	}{
		name:"xxb",
		age:90,
	}
	fmt.Println(s2)
	//anonymous  column
	w2 := worker{"sad",12}
	fmt.Println(w2)
	fmt.Println(w2.string,w2.int) //sad 12 默认字段为数据类型的名字

}
type worker struct { //anonymouse column     默认字段为数据类型的名字   匿名字段类型不能重复否则会有冲突
	 string
	 int
}

type student struct {
	name string
	age int
}