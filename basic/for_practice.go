package main

import "fmt"

func main() {
	/*
		for i:=58;i>=23;i--{
			fmt.Println(i)
		}
		//打印58-23的
	*/

	/*
		sum := 0
		for i := 0; i <= 100; i++ {
			sum += i //total = total + i
		}
		fmt.Println(sum)
	*/
	/*
		var a int = 0
		for i := 1; i <= 100; i++ {
			if i%3 == 0 && i%5 != 0 {
				fmt.Print(i,"\t")
				a++
				if a % 5 == 0{
					fmt.Println()
				}
			}

		}
		fmt.Println(a)

	*/
	for i := 1; i <= 5; i++ {
		for i := 1; i <= 5; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i;j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	for j:=1;j<=3;j++{
		fmt.Printf("%d x %d = %d\t", 3, j, 3*j)
	}

}
