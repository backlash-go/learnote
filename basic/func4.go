package main

import "fmt"

func main() {
	/*
		多返回值   顺序 个数 类型 要一一对应
	return  函数调用返回值， 结束函数后面的代码不再执行， 一般放在函数的最后一行
	空白标志符，用于舍弃数据的， 不可读


	函数 RETURE语句的使用
	return : 函数中有分支， 循环， 要保证无论执行那个分支都有return语句被执行
	函数中没有定义返回值， 可以只用return 用于专门结束函数的执行

	*/
	res1, res2 := rectangel(5, 3)
	fmt.Println(res1, res2)
	res3, res4 := rectangel2(5, 3)
	fmt.Println(res3, res4)
	_,res5 := rectangel2(5,3)
	fmt.Println(res5)
	fmt.Println(f1(30))
	f2()
}

//求矩形的周长与面积
func rectangel(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area

}

func rectangel2(len, wid float64) (peri float64, are float64) {
	peri = (len + wid) * 2
	are = len * wid
	return
}

func f1(age int)(int)  {
	if age >= 0 {
		return age
	}/*else {
		return 0
	}*/
	fmt.Println(0)
	return 0

}

func f2()  {
	for i:=0;i<=5 ;i++  {
		if i==5 {
			//break   强制结束循环
			return    //结束函数

		}
		fmt.Println(i)

	}
	fmt.Println("f2 is over")
}