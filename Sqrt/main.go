// WAP for getting a square root of a given number
package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64
	fmt.Println("Enter the given number :")
	fmt.Scan(&n)
	result := math.Sqrt(n)
	fmt.Println("Square root of given number is :", result)
}
