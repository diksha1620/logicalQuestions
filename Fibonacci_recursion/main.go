// WAP to print Fibonacci series with recursion

package main

import "fmt"

func printFibonacii(n int) {
	var n3, n1, n2 int = 0, 0, 1

	for i := 1; i <= n; i++ {

		fmt.Print(" ", n1)

		n3 = n1 + n2

		n1 = n2

		n2 = n3

	}
}

func main() {
	printFibonacii(10)
}
