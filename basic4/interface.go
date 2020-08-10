package main

import "fmt"

func main() {
	/*
	接口类型声明中的这些方法所代表的就是该接口的方法集合
	对于任何数据类型只要它的方法集合中完全包含了一个接口的全部特征（即全部的方法），那么它就一定是这个接口的实现类型


	接口：是一组方法的签名
	 	 当某个类型为这个接口的所有方法提供了方法的实现， 它被称为实现接口
		接口对象不能访问实现类中的属性

	多态，一种事物的多种形态  go语言通过接口模拟
		就一个接口的实现
			1 看成实现本身的类型， 能够访问实现类的方法与属性
			2 看成接口类型，那就只能访问接口中的方法
	 */

	m := mouse{name:"逻辑鼠标"}
	fmt.Println(m.name)
	f := flaskdisk{name:"U盘"}
	fmt.Println(f.name)

	test(m)
	test(f)
	var us usb
	us = m
	us.end();us.start()
	fmt.Println("_____++++++")
	m.end();m.start()
	//接口对象不能访问实现类中的属性 fmt.Println(us.name)
}


//定义接口

type usb interface {
	start()
	end()
}

//实现类
type mouse struct {
	name string
}

type flaskdisk struct {
	name string
}

func (m mouse) start()  {
	fmt.Println(m.name,"mouse is start")

}
func (m mouse) end()  {
	fmt.Println(m.name,"mouse is end")

}

func (f flaskdisk) start()  {
	fmt.Println(f.name,"flaskdisk is start")

}
func (f flaskdisk) end()  {
	fmt.Println(f.name,"flaskdisk is end")

}

//测试方法
func test(u usb)  {   //u = m   u = f

	u.start()
	u.end()
}