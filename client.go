package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp","127.0.0.1:12345")
	if err!= nil{
		fmt.Println("dial ")
		return
	}
	defer conn.Close()
for {
	//主动发送数据
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println("dial ")
		return
	}
	fmt.Println(n)
	time.Sleep(time.Second)
}

}
