package main

import "fmt"

func send(out chan<- int)  { //写
	out <- 789
}

func recevie(in <-chan int)  {  //读
	num := <- in
	fmt.Println(num)
}

func recevie2(in <-chan int)  {  //读
	num2 := <- in
	fmt.Println(num2)
}

func main() {

	ch := make(chan int)

	go send(ch)
	go recevie(ch)
	go recevie2(ch)

	for  {
		;
	}
}
