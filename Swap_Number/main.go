// Swap Values without using any variable a=10, b=12

package main

import "fmt"

func main() {
	a := 10
	b := 12

	// b = a + b
	// a = b - a
	// b = b - a
	a, b = b, a
	fmt.Printf("a=%v \nb=%v", a, b)

}
