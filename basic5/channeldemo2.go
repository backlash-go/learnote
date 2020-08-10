package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int,3)
	go func() {
		for i:=0;i<5;i++{
			chan1 <- i

			fmt.Println("sub ",i)
		}
		fmt.Println("finish")

	}()

	time.Sleep(time.Second * 2)

	for i:=0;i<5;i++{
		num := <- chan1
		//time.Sleep(time.Millisecond * 3)
		fmt.Println("patrent", num)

	}
}
