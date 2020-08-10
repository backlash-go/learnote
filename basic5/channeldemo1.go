package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan bool)  //控制打印输出的 顺序 先写 再读
	go func() {
		for i:=0;i<5;i++{
			chan1 <- i

			//time.Sleep(time.Millisecond * 300)
			fmt.Println("sub ",i)
			chan2 <- true
		}
		fmt.Println("finish")

	}()

	//time.Sleep(time.Second * 2)

	for i:=0;i<5;i++{
		num := <- chan1
		//time.Sleep(time.Millisecond * 3)
		<- chan2
		fmt.Println("patrent", num)

	}
}
