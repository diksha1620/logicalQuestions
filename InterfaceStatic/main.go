package main

import "fmt"

type tank interface {
	Area() float64
	Volume() float64
}
type values struct {
	radius float64
	height float64
}

func (m values) Area() float64 {
	return 3.14 * m.radius * m.radius

}

func (m values) Volume() float64 {
	return 2 * 3.14 * m.height * m.radius
}
func main() {
	var t tank
	t = values{11, 22}
	fmt.Println("Area: ", t.Area())
	fmt.Println("Area: ", t.Volume())

}
