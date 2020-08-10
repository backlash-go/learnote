package main

import (
	"fmt"
	"sort"
)

func main() {
	//创建map, slice map 默认是为 nil， **》》map声明没有初始化不能使用（nil map）
	var map1 map[int]string
	var map2 = make(map[int]string) //make直接创建map可以使用
	var map3 = map[string]int{"a": 1,"b":2}
	//map1[1] = "a"
	map2[2] = "b"
	fmt.Println(map3, map2, map1)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	//map使用
	map1[1] = "aa"
	map1[2] = "bb"
	map1[3] = "cc"
	map1[4] = "dd"
	map1[5] = "aa"
	fmt.Println(map1, map1[1])
	fmt.Println(map1[22]) //key不存在获取的是value数据类型的默认值
	//
	v1, ok := map1[22]
	if ok {
		fmt.Println("yes", v1)

	} else {
		fmt.Println("no v1", v1)
	}
	//修改数据
	map1[1] = "xixianbin"
	fmt.Println(map1[1])
	//删除数据
	delete(map1, 3)
	fmt.Println(map1)
	delete(map1, 30) //key不存在时没有影响
	fmt.Println(len(map1))

	//map的遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	//先把key用切片存储，然后排序切片，通过for_range遍历
	map1[3] = "saasd"
	keys := make([]int, 0, len(map1))
	for k, _ := range map1 {
		fmt.Println(k)
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s1 := []string{"Apple", "Window", "Orange", "abc", "王二狗", "acd"} //按照编码值排序， 大写字母A为32 小写32+65
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)


	//map结合slice使用
	var m1 = make(map[string]string)
	m1["name"] = "xxb"
	var m2 = make(map[string]string)
	m2["name"] = "lyq"
 	var m3 = make(map[string]string)
 	m3["name"] = "xly"
fmt.Println(m1,m2,m3)


 	s11 := make([]map[string]string ,0,3)
	s11 = append(s11,m1)
	s11 = append(s11,m2)
	s11 = append(s11,m3)
	for  i,v := range s11{
		fmt.Println(i+1,v["name"],v)
	}

	//map是引用类型的数据
    map5 := make(map[string]map[string]string)
    map6 := make(map[string]string)
    map6["a"] = "aa"
    map5["A"] = map6
    fmt.Println(map5,map6,map5["A"]["a"])


    //map 嵌套，
	var mia = map[string]map[string]string{
		"chant_stu": {"MINI_APP_ID": "asad", "MINI_APP_SECRET": "b", "MINI_APP_TOKEN": "c"},
		"chant_tea": {"MINI_APP_ID": "1", "MINI_APP_SECRET": "2", "MINI_APP_TOKEN": "3"}}

	//fmt.Println(mia["chant_stu"]["MINI_APP_ID"])
	//

	//if aa1,ok:=mia["chants_stu"]; ok {
	//	//fmt.Println("aa")
	//	fmt.Println(ok)
	//	fmt.Println(aa1)
	//}
	//fmt.Println(ok)


}
