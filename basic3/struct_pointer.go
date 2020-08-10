package main

import "fmt"

func main() {
	/*

		数据类型
			值类型    int float bool string array  struct
			引用类型  slice map function pointer
		结构体指针  通过指针可以变为引用类型  一般通过new函数创建指针,
	   new  不是nil

	*/

	p1 := personer{}
	p1.name = "王二狗"
	p1.age = 20
	p1.sex = "男"
	p1.address = "杭州"
	fmt.Println(p1)
	fmt.Printf("%T,%p\n", p1, &p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%T,%p\n", p2, &p2)
	p2.name = "asd"

	fmt.Println(p1, p2) //{王二狗 20 男 杭州} {asd 20 男 杭州}  //值传递  是深拷贝

	////定义结构体的指针     引用传递  是浅拷贝
	//
	var pp1 *personer
	pp1 = &p1
	fmt.Printf("%T,%p\n", pp1, &pp1)
	fmt.Println(*pp1)
	p1.name = "struct"

	fmt.Println(p1, *pp1, p2)
	//
	////使用内置函数new 专门创建某种类型的指针的函数
	//
	//pp2 := new(personer)
	//fmt.Printf("%T,%p", pp2, pp2)
	//fmt.Println(*pp2)
	//(*pp2).name = "pp22"
	//(*pp2).age = 535         //标准写法
	//pp2.address = "sadfafasf"  //简写赋值
	//fmt.Println(*pp2)
	//
	////new
	//pp3 := new(int)
	//fmt.Println(pp3,*pp3)//0xc00001a078 0


}

type personer struct {
	name    string
	age     int
	sex     string
	address string
}
