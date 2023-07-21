//     1
//   2 3 4
// 5 6 7 8 9
// 9 8 7 6 5
//   4 3 2
//     1

package main

import "fmt"

func main() {
	num := 1
	for i := 1; i <= 3; i++ {
		for j := 3; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Printf("%d", num)
			num++
		}
		fmt.Println()
	}
	num2 := 9
	for i := 3; i >= 1; i-- {
		for j := 3; j > i; j-- {
			fmt.Print(" ")
		}
		for k := i*2 - 1; k >= 1; k-- {
			fmt.Printf("%d", num2)
			num2--
		}
		fmt.Println()
	}
}
