package main

import (
	"fmt"
)

func main() {

	p1 := persone{"王二狗",30}
	fmt.Println(p1.name,p1.age)
	p1.eat()
	fmt.Printf("%T\n",p1)

	s1 := stu{persone{"xxb",30},"qingfu"}
	fmt.Println(s1.name)  //s1.persone.name
	s1.eat()
	s1.study()

}


//匿名字段包含一个方法， 那个包含该匿名字段的结构体也可以调用该方法

type persone struct {
	name string
	age int
}

type stu struct {
	persone
	school string
}

func (p persone) eat(){
	fmt.Println("parent class")
}

func (s stu) study(){
	fmt.Println("study")
}

func (s stu) eat(){
	fmt.Println("sub class")
}