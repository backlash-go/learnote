package main

import (
	"errors"
	"fmt"
)

type Printer func(content string) (n int, err error) //定义一个函数类型

type opertate func(x, y int) int

func printTostd(contents string) (bytenum int, err error) {
	return fmt.Println(contents)

}



func calculate(x, y int, op opertate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func main() {
	op := func(x, y int) int { return x + y}

	var p Printer  //声明一个函数类型的变量
	p = printTostd //引用 printTostd函数
	p("something")
	fmt.Println(calculate(4,5, op))
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)

	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
}
