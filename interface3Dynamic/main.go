package main

import "fmt"

type user interface {
	Area() float64
	Volume() float64
}

func main() {
	var t user
	fmt.Printf("Type : %T", t)
}
