// 5 4 3 2 1
// 5 4 3 2
// 5 4 3
// 5 4
// 5

package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		for j := 5; j >= 5-i+1; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
