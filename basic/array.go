package main

import "fmt"

func main() {
	var arr1 [4] int // var arr1 = [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println("", cap(arr1)) //数组能够存储的最大容量CAPACITY
	fmt.Println("", len(arr1)) //数组实际存储的数据量
	arr1[0] = 100
	fmt.Println(arr1[0])
	// 数组定长， 长度与容量相同

	var arr2 [4] int //[0,0,0,0]
	fmt.Println(arr2)

	var b = [4]int{1, 2, 3} //[1,2,3,0]
	fmt.Println(b)

	var c = [4]int{1: 1, 3: 2} //[0 1 0 2]
	fmt.Println(c)

	var e = [5]string{1: "sd", 4: "qa"}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(f))
	fmt.Printf("%p\n", &f)

	g := [...]int{1: 3, 9: 3}
	fmt.Println(len(g))
	fmt.Printf("%p\n", &g)
	//数组的遍历
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	for index, value := range f {
		fmt.Println(index, value)
	}

	for _, v := range f {
		fmt.Println(v)
	}
	//数组是值类型 size(type)
	arr3 := [4]float64{1.23, 4.34, 3.14}
	arr4 := [3]string{"a", "v", "s"}
	fmt.Printf("%T,%T,%T\n", f, arr3, arr4)

	num := 10
	num1 := num //值传递
	num1 = 20
	fmt.Printf("%p,%p\n", &num, &num1)
	fmt.Println(num, num1)
	//数组赋值
	arr5 := arr1 //数组也是值传递
	fmt.Println(arr1)
	fmt.Println(arr5)

	arr5[0] = 200
	fmt.Println(arr1)
	fmt.Println(arr5)

	arr6 := [5]int{15, 23, 8, 15, 8}

	for i := 1; i < 5; i++ {
		for j := 0; j < len(arr6)-i; j++ {
			if arr6[j] > arr6[j+1] {
				arr6[j], arr6[j+1] = arr6[j+1], arr6[j]
			}
		}
		fmt.Println(arr6)

	}
	//遍历二维数组
	arr7 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(arr7[0], arr7[0][3], len(arr7), len(arr7[0]))
	fmt.Printf("%p\n", &arr7)

	for i := 0; i < len(arr7); i++ {
		for j := 0; j < len(arr7[i]); j++ {
			fmt.Print(arr7[i][j], "\t")
		}
		fmt.Println()

	}

	for _, arr := range arr7 {
		for _, arrr := range arr {
			fmt.Print(arrr, "\t")
		}
		fmt.Println()
	}
}
