package main

import (
	"fmt"
)

func main() {
	switch { //省略默认是BOOL类型
	case true:
		fmt.Println("1")
	case false:
		fmt.Println("0")

	}

	score := 88
	switch {
	case score < 60 && score >= 0:
		fmt.Println("bad")
	case score >= 60 && score <= 100:
		fmt.Println("good")
	default:
		fmt.Println("input is false")

	}

	//var letter string = "c"
	switch letter := "c"; letter {
	case "A", "B", "C":
		fmt.Println("a")
	case "c", "d":
		fmt.Println("b")

	}

	day := 0
	month := 5
	year := 2019
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%4 == 0 && year%100 == 0 || year%400 != 0 {
			day = 28
		} else {
			day = 29
		}
	}
	fmt.Printf("day is %d\n", day)
}
