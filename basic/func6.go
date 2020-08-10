package main

import "fmt"

func main() {
	/*
			递归函数  一个函数自己调用自己
		求和 1-5
		r1(5)       r1(n-1) + n
		  r1(4) + 5
			r1(3) + 4
		     r1(2) + 3
		       r1(1)=1
	*/
	t := r1(5)
	fmt.Println(t)
	fmt.Println(fi(7))
}

func r1(n int) int {
	if n == 1 {
		return 1
	}
	return r1(n-1) + n

}

//斐波那契数列
//项 1 2 3 4 5 6  7  8  9 。。。
//值 1 1 2 3 5 8 13 21  34 。。。
func fi(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fi(n-1) + fi(n-2)

}
