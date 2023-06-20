package main

import "fmt"

func main() {
	var n, i int

	fmt.Print("Enter the Even Array Size = ")
	fmt.Scan(&n)

	evenArr := make([]int, n)

	fmt.Print("Enter the Even Array Items  = ")
	for i = 0; i < n; i++ {
		fmt.Scan(&evenArr[i])
	}
	fmt.Print("\nThe List of Even Numbers in evenArr = ")
	for i = 0; i < n; i++ {
		if evenArr[i]%2 == 0 {
			fmt.Print(evenArr[i], " ")
		}
	}
	fmt.Println()
}
