package main

import "fmt"

func main() {
	/*
		goto 语句

	*/
	var a int = 10
	LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("%d\n", a)
		a++
	}
}
