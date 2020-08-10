package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"math/rand"

)


func main() {
	//var ma1 = make(map[int]int)
	//ma1[1] = 1
	//
	//a := ma1[2]
	//fmt.Println(ma1[1],a)
	//
	//
	//t := <- time.After(time.Second * 3)
	//fmt.Println(t)
	//map1 := make(map[string]interface{})
	//map1{"msg":"sa","code",123}
	//fmt.Println(map1)

	//fmt.Println(map[string]interface{}{"msg":"sa","code":123}["msg"])
	//filename := path.Join("/Users/xixianbin/go/src/Project1/basic5/file.txt")

	//fmt.Println(filename)
	//logfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	panic(fmt.Errorf("failed create log file, err is %s", err))
	//}
	//if _, err := logfile.Write([]byte("appended some data\n")); err != nil {
	//	logfile.Close() // ignore error; Write error takes precedence
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%T\n", logfile)
	//s := []byte("aaa")
	//fmt.Println(s)

	//slice1 := make([]string ,3,8)
	//fmt.Printf("%T,%p,%p,%v\n",slice1,&slice1,slice1,slice1)
	//slice1 = append(slice1,"a","b")
	//fmt.Printf("%T,%p,%p,%v\n",slice1,&slice1,slice1,slice1)
	//
	//slice1 = append(slice1,"a","b","c","d","e","f","s")
	//fmt.Printf("%T,%p,%p,%v\n",slice1,&slice1,slice1,slice1)
	//
	//slice2 := slice1
	//fmt.Printf("%p,%p\n",&slice2,slice2)

	//slice1 = append(slice1,"a","b")
	//fmt.Printf("%T,%p\n",slice1,&slice1)
	//fmt.Println(slice1,len(slice1),cap(slice1))
	//
	//
	//slice1 = append(slice1,"a","b","c","d","f","s")
	//fmt.Printf("%T,%p\n",slice1,&slice1)
	//
	//
	////slice2 := append(slice1, "a","v","c")
	////fmt.Printf("%T,%p,%v\n",slice2,slice2,slice2[1:])
	//
	//fmt.Println(slice1,len(slice1),cap(slice1))
	//
	//
	//slice2 := append(slice1,"a","b","c","d","f","s")
	//fmt.Printf("%T,%p,%p\n",slice2,slice2,&slice2)
	//
	//fmt.Println(len(slice2),cap(slice2))


	//test := []struct{
	//	add string
	//	pwd string
	//	choiceDB string}{
	//		{"1","2","3"},
	//		{"4","5","6"},
	//}
	//fmt.Printf("%v\n",test)
	//fmt.Printf("%T\n",test)

	//var s []string
	//

	//g := Args{}.Add("id")
	//fmt.Printf("%T\n",g)
	//fmt.Println(g)
	//
	//}
	//type Args []interface{}
	//
	//// Add returns the result of appending value to args.
	//func(args Args) Add(value ...interface {}) Args {
	//	return append(args, value...)
	//slice1 := []string{"a","b","c","d","f"}
	//slice2 := []string{"1","2","3"}
	//slice1 = append(slice1, slice2...)
	//fmt.Println(slice1)
	//
	//
	//type u struct {
	//	a string
	//	b string
	//}
	//
	//st1 := u{}
	//
	//fmt.Printf("%T,%s\n",st1,st1)

	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"foo": "bar",
	//	"nbf": time.Date(215, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	//})
	//
	//// Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString([]byte("abcd"))
	//if err!= nil{
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(tokenString)
	//	h := sha256.New()
	//	io.WriteString(h,"13960679708")
	//	io.WriteString(h,"xxb@qq.cpm")
	//
	//	token := fmt.Sprintf("%x", h.Sum(nil))
	//	fmt.Println(token)
	//	fmt.Println("200d80623315a8796cc36534f8eeb9c59b9a64e50a21c2366e3ee97393030ce1")
	//    rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Int())
	//aaa := strconv.Itoa(rand.Int())
	//fmt.Println(aaa)
	//
	//
	//strings.HasPrefix("/heath","/heath")
	//
	//fmt.Println(strings.HasPrefix("/heath","/heath"))
	//    ch1 := make(chan int,3)
	//    ch1 <- 2
	//    ch1 <- 1
	//    ch1 <- 3
	//    elem1, bl := <- ch1
	//    fmt.Println(elem1,bl)
	//	close(ch1)
	//    close(ch1)
	//	elem2,bl2:= <- ch1
	//	fmt.Println(elem2,bl2)

	//	fmt.Println(elem1,elem2,bl)
	//fmt.Printf("%T,%v\n",ch1,ch1)
	//
	//ch2 := make(chan int)
	//fmt.Printf("%T", ch2)
	//    fmt.Println(ch2)

	//var ch3 chan int
	//ch3 <- 4
	//el3 := <- ch3
	//fmt.Println(ch3,el3)
	// var aa string = "ss"
	// fmt.Printf("%T,%p\n",aa,&aa)
	// var b string = aa
	//fmt.Printf("%T,%p\n",b,&b)
	//
	// var ss1 = []int{1, 2, 3, 4}
	// var ss2 []int = ss1
	// fmt.Printf("%p\n",&ss1)
	// fmt.Printf("%p\n",&ss2)
	//
	// ss1 = append(ss1,5,6,7,8)
	//fmt.Printf("%p,%p\n",ss1,&ss1)

	//channel1 := make(chan <- int)
	//fmt.Printf("%T\n",channel1)
	//
	//channel2 := make(<- chan int )
	//fmt.Printf("%T\n",channel2)

	//

	//

	//defer func(){
	//	fmt.Println("Enter defer function.")
	//	if p := recover(); p != nil {
	//		fmt.Printf("panic: %s\n", p)
	//	}
	//	fmt.Println("Exit defer function.")
	//}()
	//// 引发panic。
	//panic(errors.New("something wrong"))
	//fmt.Println("Exit function main.")


	type Message struct {
		Name string
		Body string `json:"body"`
		Time int64
	}

	me1 := Message{"XXB",`{"Name":"XXB","Body":"sex","Time":12}`,12}
	b, err := json.Marshal(me1)
	b1 := string(b)
	fmt.Println(err)
	fmt.Printf("%T,%v\n",b1,b1)

}

func test() string{
	//defer fmt.Println("a")
	//defer fmt.Println("b")
	panic(errors.New("this is a test panic"))
	return "rr"


}



