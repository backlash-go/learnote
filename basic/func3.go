package main

import "fmt"

func main() {
	/*
	 函数的返回值
	*/
	s1 := []int{1, 2, 3}
	nedSum(s1...)
	fmt.Println(getSum2())
}

func nedSum(nums ...int) int { //num 是个切片 []int
	fmt.Printf("%T\n", nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
	return sum

}

func getSum2() (sum int) { // 0 1
	fmt.Println(sum)
	return 1
}
