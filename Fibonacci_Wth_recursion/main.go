package main

import "fmt"

func main() {
	var n3, n1, n2 int = 0, 0, 1
	n := 10
	for i := 1; i <= n; i++ {

		fmt.Print(" ", n1)

		n3 = n1 + n2

		n1 = n2

		n2 = n3

	}
}
