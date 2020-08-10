package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	var s1 *cc
	s1 = &cc{}

	s1 = &cc{1, 1, "v"}
	fmt.Println(s1)
	fmt.Println(s1.b)
	fmt.Println(s1.d)

	fmt.Printf("%T,%s\n", s1.d, s1.d)
	uid := fmt.Sprintf("%s", uuid.New())
	iid := uuid.New().String()
	fmt.Println(iid)
	fmt.Println("+++++")
	fmt.Printf("UUIDv4:%T, %s,%s\n", iid, uid, len(uid))

}

type cc struct {
	a int
	b int
	d string
}

type bb struct {
	ab int
	bb int
	db string
}
