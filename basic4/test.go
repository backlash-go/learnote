package main

import (
	"fmt"
	"time"
)

func printer(str string)  {
	for _, ch := range str{
		fmt.Printf("%c",ch)
		time.Sleep(time.Millisecond * 200)
	}

}

func main() {
	str := "strin"
	for _, ch := range str{
		//fmt.Printf("%c",ch)
		fmt.Println(ch)
		time.Sleep(time.Millisecond * 200)
	}
}