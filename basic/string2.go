package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "helloworld"
	fmt.Println(strings.Contains(s1, "abcg")) //是否包含指定的内容
	fmt.Println(strings.ContainsAny(s1,"abcd"))//只要有一个字符在S1出现过就是TRUE
	fmt.Println(strings.Count(s1,"llo")) //计数
	fmt.Println(strings.LastIndex(s1,"l")) //查找最后一次出现的
	fmt.Println(strings.Index(s1,"d"))  //查找substr在s1中的第一个位置// 不存在则返回-1
	fmt.Println(strings.IndexAny(s1,"abcdh")) //查找任意一个字符在s1中的位置，同时出现返回0
	fmt.Println(strings.IndexAny(s1,"abcd"))

	s2:="2019课堂笔记.txt"
	if strings.HasPrefix(s2,"2019"){  //前缀开头
		fmt.Println("yes")
	}

	if strings.HasSuffix(s2,".txt") {  //结尾
		fmt.Println("是个文本文件")

	}

	//字符串拼接
	ss1 := []string{"as","fs","saf"}
	s3 := strings.Join(ss1,"*")
	fmt.Println(s3)

	//切割
	s4 := "123,535,sdf,fds"
	ss2 := strings.Split(s4,",")
	fmt.Println(ss2,ss2[0])
	//重复字符串 自己拼接自己五次
	s5 := strings.Repeat("saf",5)
	fmt.Println(s5)

	s6 := strings.Replace(s1,"l","*",2)//he**oworld
	s7 := strings.Replace(s1,"l","*",-1)//he**owor*d	s6 := strings.Replace(s1,"l","*",3)//he**owor*d 替换所有

	fmt.Println(s6, s7)

	s8 := "AFSksfL12Lf"
	fmt.Println(strings.ToLower(s8))
	fmt.Println(strings.ToUpper(s8))

	s9 := s8[0:3]
	fmt.Println(s9)

}
