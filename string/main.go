package main

import "fmt"

func main() {
	s := []string{"apple", "bob", "pineapple", "zoz", "tomato", "grapes", "ant", "banana", "boy", "cat", "dog"}
	var res []string

	for _, str := range s {
		if Palindrome(str) {
			res = append(res, str)
		}
	}

	fmt.Println(res)
	fmt.Println(s)
}

func Palindrome(str string) bool {
	lastIdx := len(str) - 1

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[lastIdx-i] {
			return false
		}
	}

	return true
}
