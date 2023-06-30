package main

import "fmt"

func main() {
	fmt.Println("Pattern problem for pyramid")
	n := 15

	for i := 1; i <= (n/2 + 1); i++ {
		for j := 1; j <= (n - i); j++ {
			fmt.Print(" ")
		}
		var skip int = 0
		for j := 1; j <= i*2-1; j++ {
			if i > 2 && i <= n/2 && (j == 1 || j == i || j == (i*2)-1) {
				fmt.Print("*")
				skip = i - 2
				continue
			}
			if i > 2 && i <= n/2 && skip > 0 {
				fmt.Print(" ")
				skip--
				continue
			}
			fmt.Print("*")
		}

		fmt.Println()
	}
	for i := n / 2; i > 0; i-- {
		for j := 1; j <= (n - i); j++ {
			fmt.Print(" ")
		}
		var skip int = 0
		for j := 1; j <= i*2-1; j++ {
			if i > 2 && i <= n/2 && (j == 1 || j == i || j == (i*2)-1) {
				fmt.Print("*")
				skip = i - 2
				continue
			}
			if i > 2 && i <= n/2 && skip > 0 {
				fmt.Print(" ")
				skip--
				continue
			}
			fmt.Print("*")
		}

		fmt.Println()
	}
}

// for i := n - 2; i > 0; i-- {
// 	for j := i; j < n-1; j++ {
// 		fmt.Print(" ")
// 	}
// 	for j := 1; j <= (2*i - 1); j++ {
// 		fmt.Print("*")
// 	}
// 	fmt.Println()

// 	}
// }
