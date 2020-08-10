package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	注意事项 ， 一个channel 只能在两个go程中使用，
				一个channel 在一个go程中进行读写 会发生deadlock死锁
	 */
	chan1 := make(chan string,8)
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println("zi go 程， i =", i)

			time.Sleep(time.Second)
		}
		chan1 <- "avc"
		chan1 <- "bavc"
		chan1 <- "cavc"
		chan1 <- "davc"
		chan1 <- "favc"
	}()

	str1 := <- chan1
	str2 := <- chan1

	fmt.Println(str1,str2,len(chan1),cap(chan1)) //len(channel) 缓冲区剩余的数据个数， 无缓冲channel的len始终为0  //avc bavc 3 8

}
