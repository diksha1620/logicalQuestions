// Given number is prime or not

package main

import "fmt"

func main() {
	var n int = 11
	var count int = 0
	if n < 2 {
		fmt.Printf("%v is not prime number", n)
	}
	for i := 1; i < n; i++ {
		if n/i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Printf("%v is not prime number", n)
	} else {
		fmt.Printf("%v is prime number", n)
	}
}
