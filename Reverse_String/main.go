// Find the reverse of the string

package main

import "fmt"

func main() {

	str := "ReverseString"
	strRev := ""
	for _, v := range str {
		strRev = string(v) + strRev
	}

	fmt.Println(str)
	fmt.Println(strRev)
}
