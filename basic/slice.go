package main

import "fmt"

func main() {
	/*
	   1 切片引用了一个底层的数组
	   2 自动扩容 超过容量扩容为原来的两倍
	   3 一旦扩容重新指向一个新的数组
	   4 切片不存储数据，所以修改数据其实就是修改底层的数组的值
	*/
	var s1 [] int
	fmt.Println(s1)

	var s2 = []int{1, 2, 3, 4}
	fmt.Printf("%T,%T\n", s1, s2)

	s3 := make([]int, 4, 8)
	fmt.Printf("%d,%d\n", cap(s3), len(s3))

	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Printf("%p\n",s3)

	s3 = append(s3, 4,5,6,7,8)
	fmt.Printf("%p\n",s3)
	s3 = append(s3, s2...)
	fmt.Println(s3[0], s3[1], s3[2], s3[3], s3[4], s3)

/*
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	for k, v := range s3 {
		fmt.Print(k, v)
		fmt.Println()
	}

	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	//切片扩容会生成一个新的数组
	s4 := arr[0:3] //[0:3)
	s5 := arr[:5]
	s6 := arr[:]
	s7 := arr[5:]
	fmt.Printf("%d,%T\n", s4, s4)
	fmt.Println(s5, s6)
	fmt.Printf("%d,%d\n", len(s4), cap(s4))
	fmt.Printf("%d,%d\n", len(s7), cap(s7))

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	fmt.Printf("%d,%d,%p,%p\n", arr1, arr2, &arr1, &arr2)

	ss1 := []int{4, 5, 6}
	ss2 := ss1
	fmt.Printf("%T,%d,%d,%p,%p\n", ss1, ss1, ss2, ss1, ss2)
	//实现切片的深拷贝

	sl1 := []int{1, 2, 3}
	sl2 := make([]int, 4)
	for i := 0; i < len(sl1); i++ {
		sl2 = append(sl2, sl1[i])

	}
	//fmt.Println(sl1, sl2)
	fmt.Printf("%d,%d,%p,%p\n",sl1,sl2,sl1,sl2)
	sl1[0] = 100
	fmt.Println(sl1,sl2)
	//copy
	sl3 := []int{1,2,3,4}
	sl4 := make([]int,4)
	sl5 := []int{11,12,13,14}
	//copy(sl3,sl5)  //将sl5 拷贝到sl3
	//copy(sl3,sl5[0:2]) 可以指定拷贝长度
	copy(sl3[1:],sl5[0:2])
	fmt.Println(sl3,sl5,sl4)
*/

}


