package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {

			fmt.Println("这是个匿名函数。。。")
			//time.Sleep(time.Millisecond * 200)
		}

	}()

	for i := 0; i < 5; i++ {
		fmt.Println("main...")
		runtime.Gosched()  //主动出让CPU使用权一次，给其他go程增加获取CPU的机会
	}
}
