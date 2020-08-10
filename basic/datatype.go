package main

import "fmt"

func main() {
	/*
		go 语言数据类型
		1，基本数据类型
			布尔类型： bool    true flase

			数值类型:
				整数 int，uint8，16，32，64
		 			byte：uint8
					rune：int32
				浮点数
					float 32  float 64
				复数
			字符串类型
		2，复合数据类型
	*/

	var i3 int //根据计算32位 64位。默认int多少位
	i3 = 1000
	var i4 int64
	//i4 = i3   //语法角度int  int64 不认为是同一类型不能赋值
	fmt.Println(i3, i4) //./datatype.go:22:5: cannot use i3 (type int) as type int64 in assignment
	var i5 byte
	i5 = 0001
	fmt.Println(i5)
	//byte 赋值 互换uint8
	var i6 uint8
	i6 = 100
	var i7 byte
	i7 = i6
	fmt.Println(i6, i7)

}
