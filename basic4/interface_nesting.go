package main

import "fmt"

func main() {
	/*
		interface nesting
	*/
	var c1 Cat = Cat{}
	c1.test1()
	c1.test2()
	c1.test3()
	fmt.Println("----------------------")
	var a1 A = c1 //a1 only use test1 method
	a1.test1()
	fmt.Println("----------------------")
	var b B = c1
	b.test2()

	fmt.Println("----------------------")
    var c C = c1
    c.test1();c.test2();c.test3()
	fmt.Println("----------------------")
    var c2 A = a1
    c2.test1()
}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {
}

func (c Cat) test1() {
	fmt.Println("test1...")
}
func (c Cat) test2() {
	fmt.Println("test2...")
}
func (c Cat) test3() {
	fmt.Println("test3...")
}
