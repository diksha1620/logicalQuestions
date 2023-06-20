// Code to find even numbers in a list
package main

import "fmt"

func main() {
	n := 22

	if n%2 == 0 {
		fmt.Printf("%v is even number", n)
	} else {
		fmt.Printf("%v is odd number", n)
	}
}
