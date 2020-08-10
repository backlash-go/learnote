package main

import "fmt"

func main() {


	//别名使用
	type myint = int  //别名
	var a myint
	b := 100
	a = b
	fmt.Println(a)
	fmt.Printf("%T,%T\n",a,b)
}
