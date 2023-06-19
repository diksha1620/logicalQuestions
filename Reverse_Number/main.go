// WAP to reverse an integer without converting it to a string,
// without using any builtin methods.

package main

import "fmt"

func main() {
	var number int = 12345
	var rev int = 0
	for number > 0 {
		rem := number % 10
		rev = rev*10 + rem
		number = number / 10
	}
	fmt.Println("The reverse number is: ", rev)
}
