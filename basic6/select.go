package main

import (
	"fmt"
	"time"
)

func main() {
	/*

	 */
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for   {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-quit :
				//return    //返回当前函数
				//runtime.Goexit()  //结束go程
				//os.Exit(0)   结束进程
				//break   跳出当前case分支
				break
		}
		fmt.Println("over")

			
		}

	}()

	for i:=0;i<7 ;i++  {
		ch <- i
time.Sleep(time.Second * 1)
	}
	quit <- true
}
