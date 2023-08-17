package main

import "fmt"

func Name(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "Diksha"
	}
	close(c)
}

func main() {
	c := make(chan string)
	go Name(c)

	for msg := range c {
		fmt.Println(msg)

	}
}
