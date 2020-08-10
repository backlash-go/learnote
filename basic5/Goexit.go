package main

import (
	"fmt"
	"os"
	"runtime"
)

func mygo() {
	defer fmt.Println("cccc")
	//return
	//runtime.Goexit()
	os.Exit(0)  //结束当前的进程
	test()
	defer fmt.Println("fffffffff")
}

func test() {
	runtime.Goexit()
	fmt.Println("dddddddd")

}

func main() {

	fmt.Println("aaaaaaaa")
	go mygo() //创建子go程

	fmt.Println("bbbbbbbbbbb")
	for {

	}
}
