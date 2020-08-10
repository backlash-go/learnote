package main

import "fmt"

func main() {
	for i:=1;i<10;i++{
		if i >5{
			break
		}
		fmt.Println(i)

	}

	for j := 1; j < 10; j++{
		if j == 4{
			continue
		}
		fmt.Println(j)
	}

}
