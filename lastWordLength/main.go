package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Hello Worlds"

	words := strings.Split(input, " ")

	wordLen := make(map[string]int)
	if len(words) > 0 {
		for _, word := range words {
			wordLen[word] = len(word)
		}
		lastWord := words[len(words)-1]
		lenLastWord, ok := wordLen[lastWord]
		if ok {
			fmt.Println(lenLastWord)
		} else {
			fmt.Println("no")
		}
	}

}
