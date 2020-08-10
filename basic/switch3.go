package main

import "fmt"

func main() {
	fmt.Printf("a")

	switch n := 1; n {
	case 1:
		fmt.Println("1111")
		fmt.Println("1111")
		fmt.Println("1111")
		fallthrough  //穿透SWITCH
	case 2:
		fmt.Println("222")
		fmt.Println("222")
		//break //强制结束CASE语句从而结束SWIRCH
		fmt.Println("222")
	case 3:
		fmt.Println("3333")
		fmt.Println("3333")
		fmt.Println("3333")




	}
}
