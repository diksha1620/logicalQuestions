package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 1; j < 10; j++ {
				ch <- "Goroutine : " + strconv.Itoa(i)
			}
		}(i)
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-ch)
	}
}
