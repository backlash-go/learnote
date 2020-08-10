package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	go程： 结合线程与协程的优点    go兵法  channel  goroutine
	go程运行状态， 初始态， 就绪， 运行， 挂起， 终结
	go 特性  主go程结束 ， 默认会释放进程地址空间， ----》 所有的子GO程自动结束

	死循环  for { ; }
	 */
	go sing()    //main 进程变成 main主go程     与另外两个go 程同时争抢CPU资源  95% 主go先，
	go dance()

	go func (){
		for i:=1;i<3;i++ {
			fmt.Println("qqqq")
			//time.Sleep(time.Millisecond * 200)
		}
	}()
	a := time.Millisecond
	fmt.Println(a)

	  //防止主go程提前结束
	for i:=1;i<3;i++ {
		fmt.Println("我是主go程")
		//time.Sleep(time.Millisecond * 200)

	}

}

func sing()  {
	for i:=1;i<3;i++ {
		fmt.Println("sing正在唱歌")
		//time.Sleep(time.Millisecond * 200)
	}

}

func dance()  {
	for i:=1;i<3;i++ {
		fmt.Println("sing正在跳舞")
		//.Sleep(time.Millisecond * 200)

	}

}

