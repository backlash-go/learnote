package main

import (
	"fmt"
	"time"
)

var mych = make(chan int)

//hwoelrllod
func main() {
	/*
	channel 是一种数据类型， 主要用来解决go程的数据同步以及数据共享的问题     管道， 双向半双工
	go程 运行在相同的地址空间， 奉行通过通信来共享内存， 而不是共享内存来通信

	c1 := make(chan int,0) 创建一个用户通信int数据类型的CHANNEL，初始容量为0   非缓冲通道，反正 >0 则为缓冲通道
	c1 := make(chan int)  容量默认为0
	a := 2
	fmt.Printf("%p\n",&a)
	ch1 := make(chan int,3)
	ch1 <- a
	ch1 <- 1
	ch1 <- 3
	elem1 := <- ch1
	elem2 := <- ch1
	elem3 := <- ch1
	fmt.Println(elem1,elem2,elem3)
	fmt.Printf("%p,%p\n",&elem1,&elem2)
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)

	 */
	fmt.Printf("%c\n", "a")
	go person1()
	go person2()

	for   { //防止主go程 先结束
		;
	}
}
func printer(str string)  {
	for _, ch := range str{
		fmt.Printf("%c",ch)
		time.Sleep(time.Millisecond * 200)
	}
	
}

func person1()  {
	printer("hello")
	mych <- 123   //使用channel 先写mych
	
}
func person2()  {
	<-mych    //如果写端没有写就会阻塞go2   例如 elem1 := <-ch1 没有elem1 变量接收 就丢弃来自channel的值
	printer("world")

}