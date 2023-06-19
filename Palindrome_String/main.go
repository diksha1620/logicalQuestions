// WAP to find the given string is Palindrome or not

package main

import (
	"fmt"
)

func main() {

	str := "MadaM"

	// str = strings.ToLower(str)

	strRev := ""
	for _, v := range str {
		strRev = string(v) + strRev
	}

	fmt.Println(str)
	fmt.Println(strRev)
	if str == strRev {
		fmt.Println("Given String is Palindrome ")
	} else {
		fmt.Println("Givent String is not Palindrome")
	}
}
