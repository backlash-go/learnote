package main

import (
	"fmt"
	"strconv"
)

func main() {
	//bool
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%T,%t\n",b1,b1)//bool,true

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n",ss1,ss1)//string,true

	//int
	s2 := "100"
	i2, err := strconv.ParseInt(s2,10,64)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n",i2,i2)

	ss2 := strconv.FormatInt(i2,10)
	fmt.Printf("%T,%s\n",ss2,ss2)

	i3, err:= strconv.Atoi("22")  //ascii to integer
	fmt.Printf("%T,%d\n",i3,i3)

	i4 := strconv.Itoa(i3)  //integer to ascii
	fmt.Printf("%T,%s\n",i4,i4)


}
