package main

import (
	"fmt"
)

func main() {
	/*
		结构体： 是由一组一系列具有相同的数据类型或者不同的数据类型的数据结构成的数据集合
	*/

	//定义结构体

	var p1 person
	fmt.Println(p1) //{ 0  }
	//结构体访问
	//方法一    //字段可以不写， 默认为字段的数据类型的默认值
	p1.name = "王二狗"
	p1.age = 20
	p1.sex = "男"
	p1.address = "杭州"
	fmt.Println(p1, p1.sex)

	//方法二
	p2 := person{}
	p2.name = "令勇气"
	p2.age = 30
	p2.sex = "女"
	p2.address = "上海"
	fmt.Println(p2, p2.name)

	//方法三
	p3 := person{name: "s", sex: "n", age: 12, address: "wuhan"}
	fmt.Println(p3, p3.address)

	p4 := person{
		name:    "sssd",
		sex:     "n",
		age:     123,
		address: "wuhan",
	}
	fmt.Println(p4)
	//
	//方法四
	p5 := person{"李小华", 20, "nan", "sad"} //这种写法要主要结构体定义成员的顺序
	fmt.Println(p5)

}

type person struct {
	name    string
	age     int
	sex     string
	address string
}
