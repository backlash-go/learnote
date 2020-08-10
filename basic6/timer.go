package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	定时器：
	单次定时
	周期定时
	 */

	//mytime := time.NewTimer(time.Second * 3)
	//fmt.Println(time.Second * 3)
	//data := <- mytime.C
	//fmt.Println(data)
	data := <- time.After(time.Second * 3)
	fmt.Println(data)
}
