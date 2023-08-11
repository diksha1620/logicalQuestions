package main

import "fmt"

type triangle interface {
	Area() int
	Volume() int
}

type cal struct {
	length  int
	breadth int
	height  int
}

func (a cal) Area() int {
	return (a.breadth * a.length) / 2
}

func (b cal) Volume() int {
	return b.breadth * b.length * b.height
}

func main() {
	var a triangle
	a = cal{11, 13, 14}
	fmt.Println(a.Area())
	fmt.Println(a.Volume())

}
