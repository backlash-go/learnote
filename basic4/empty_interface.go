package main

import "fmt"

func main() {
	/*
		空接口的使用   不包含任何方法， 所有类型都实现了空接口,空接口可以存储任意类型的数值
	*/
	var a1 a = cat{"red"}
	var a2 a = baby{"xxb", 18}
	var a3 a = 100
	fmt.Println(a1, a2, a3)
	fmt.Printf("%T,%T,%T\n", a1, a2, a3)
	test11(a1)
	test11(a2)
	test11(a3)
	test12(a1)
	test13("a", 1, "c")
	//map 使用空接口接受任意类型
	map1 := make(map[string]interface{})
	map1["a"] = 1
	map1["b"] = "bb"
	map1["c"] = baby{"babyname", 10}
	fmt.Println(map1, map1["c"])

	//切片存储任意类型
	slice1 := make([]interface{},0,10)
	slice1 = append(slice1, "as", 1, 2)
	slice2 := slice1
fmt.Printf("%p,%p\n",&slice1,&slice2)
fmt.Println(slice1,slice2)

tslice(slice1)
}
//接受切片函数
func tslice(slice []interface{})  {
	for i:=0;i<len(slice);i++ {
		fmt.Println(slice[i])
	}
}

func test11(n a) {
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}

//空接口代表接受任意数据类型   fmt print 函数都用了空接口
func test12(n interface{}) {
	fmt.Println(n)

}

func test13(n ...interface{}) {
	fmt.Println(n)
	fmt.Printf("%T\n", n)

}

type a interface {
}

type cat struct {
	color string
}

type baby struct {
	name string
	age  int
}
