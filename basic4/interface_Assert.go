package main

func main() {

}

type Shape interface {
	peri() float64
	area() float64
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64  {
return t.a * t.b
}
