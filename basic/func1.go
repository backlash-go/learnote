package main

import "fmt"

func main() {
	/*
		函数 ： 具有特定功能的代码，
			意义： 避免重复的代码， 增强程序的扩展性
		使用： 函数的定义
			函数的调用
		函数的参数：
			形参 ：用于接受外部传入的数据的变量
			实参：函数调用时给实参赋值实际的数据
	//实参必须严格匹配形参的  顺序，个数， 类型  一一对应。

	可变参数：就是切片， 要求参数类型确定， append, print函数都是可变参数
	如果一个函数同时有可变参数， 还有其他参数， 可变参数要放到最后的位置 ，一个函数参数列表最多只能有一个可变参数
	*/
	getSum(5)
	getAdd(10, 20)
	func1(1.2, 2.1, "sd")
	gSum(1,2,3)
	sl1 := []int{1,2,3,4}
	gSum(sl1...)

}

func getSum(n int) {

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i //sum = sum + i

	}
	fmt.Println(sum)
}

//func getAdd(a int,b int){
func getAdd(a, b int) { //数据类型一致可以简写
	sum := a + b
	fmt.Println(sum)
}

func func1(a, b float64, c string) {
	fmt.Printf("%.2f,%.2f,%s\n", a, b, c)

}

func gSum(nums ... int)  {   //num 是个切片 []int
	fmt.Printf("%T\n", nums)
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	fmt.Println(sum)

}

func func2(s1,s2 string,arg float64)  {

}