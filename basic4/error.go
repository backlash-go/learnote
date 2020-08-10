package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f,err := os.Open("oop.go")
	if err != nil  {
		//log.Fatal(err)  open /txt.txt: no such file or directory
 		fmt.Println(err)
		return
	}
	fmt.Println(f.Name())
	//自己创建errors
	//1
	err1 := errors.New("自己创建玩的")
	fmt.Println(err1)
	fmt.Printf("%T\n",err1)  //*errors.errorString

	//2
	err2 := fmt.Errorf("错误信心:%d",100)
	fmt.Printf("%T,%s\n",err2,err2)


}
//设计函数 返回自己定义的错误信息

func checkage(a int) error  {
	if a < 0{
		return  errors.New("年龄不合法")
	}
	fmt.Println(a)
	return nil

}