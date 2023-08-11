package main

import "fmt"

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Define the Triangle type
type Triangle struct {
	Base   float64
	Height float64
}

// Implement the Area method for Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	// Create a Triangle instance
	triangle := Triangle{
		Base:   10,
		Height: 5,
	}

	// Use the Shape interface to calculate and print the area
	var shape Shape
	shape = triangle
	fmt.Printf("Area of the triangle: %f\n", shape.Area())

	var p Polygons = Pentagon(50)
	p.Perimeter()
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()

	// var obj Object = Pentagon(50)
	// obj.NumberOfSide()
	// var pent Pentagon = obj.(Pentagon)
	// pent.Perimeter()
}
