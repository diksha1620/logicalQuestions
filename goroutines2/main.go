// Even and odd using channels and goroutines
package main

import "fmt"

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("Even", v)
	}
}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("Odd", v)
	}
}

func main() {
	var arr = []int{1, 2, 4, 2, 55, 21, 90}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, v := range arr {
		if v%2 != 0 {
			chOdd <- v
		} else {
			chEven <- v
		}
	}
}
