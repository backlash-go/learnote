package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	a := time.Now().Unix()
	b := time.Now().UnixNano()
	rand.Seed(time.Now().Unix())
	num1 := rand.Intn(10)
	num2 := rand.Intn(10)
	num3 := rand.Float64()
	fmt.Println(num1, num2, a, b, num3)
}
