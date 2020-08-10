package main

import "fmt"

func main() {
	/*
	go  同时有方法与函数
	一个方法就是包含了接收者的函数， 接收者可以是可以是命名类型或者结构体的类型的一个值或者是指针
	 所有给定类型的方法属于该类型的方法集

	对比函数
	方法：某个类别的的行为功能，需要指定接收者调用  命名可以相同，只要接收者不同
	函数：一段独立功能的代码 可以直接调用  命名不能有冲突
	 */

	w1 := worker{"xxb",18,"n"}
	w1.work()
	w2 := &worker{"sxxb",12,"n"}
	fmt.Printf("%T,%p\n", w2,w2)
	w2.work()

	w2.work2()
	w1.work2()



}

//1 定义一个结构体

type worker struct {
	name string
	age int
	sex string
}


func (w worker) work(){  //w1 = w
	fmt.Println(w.name,"is working")
}

func (w *worker) work2(){ // w = w2    w = w1的地址
	fmt.Println(w.name,"is working")
}