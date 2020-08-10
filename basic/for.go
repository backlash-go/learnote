package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println("a")

	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	for{  //for true  死循环
		fmt.Println("j")
	}
}
