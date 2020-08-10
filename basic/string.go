package main

import "fmt"

func main() {
	/*
	字符串是一些字节的集合

	 */
	a := "hello"  //一个英文字母占一个字节  中文一般占三个字节
	b := `he`
	c := "冲占a"
	fmt.Println(a,b,c,len(a),len(b),len(c),a[0])//hello he 冲占 5 2 6  字符串长度返回的是字节的个数
	// a[0] 获取字符串中第一个字节 打印的是字节
	//遍历字符串字节
	for i:=0;i<len(a);i++{
		//fmt.Println(a[i])
		fmt.Printf("%c\t\n",a[i])
	}

	for i:=0;i<len(c);i++{
		//fmt.Println(a[i])
		fmt.Printf("%c\t\n",c[i]) //中文遍历会出问题
	}

	for _,v := range a{
		//fmt.Println(v)
		fmt.Printf("%c\n", v)
	}

	//字节转字符
	slice1 := []byte{65,66,67,68,69}
	s1 := string(slice1)
	fmt.Println(s1) //ABCDE
	//字符转字节

	s2 := "abcde"
	slice2 := []byte(s2)
	fmt.Println(slice2)

	//字符串不能修改
	//s2[1] = "S"
	//fmt.Println(s2[1])
	fmt.Println("-------------------")
	b11 := []byte("hello I am a good boy")
	fmt.Println(b11)
	b22  := make([]byte,0)
	fmt.Printf("%s,%s\n",len(b22),cap(b22))
	b23 := append(b22 ,64,65,67)
	b22 = append(b22 ,64,65,67)
	fmt.Printf("%p,%v\n",b22,b22)
	fmt.Printf("%p,%v\n",b23,b23)

}
