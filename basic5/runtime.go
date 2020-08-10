package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := runtime.GOROOT()
	fmt.Println(a,runtime.Version(),runtime.NumCPU())
}
