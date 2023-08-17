// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func fibonacci(n int, ch chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	a, b := 0, 1
// 	for i := 0; i < n; i++ {
// 		ch <- a
// 		a, b = b, a+b
// 	}
// 	close(ch)
// }

// func main() {
// 	n := 10 // Number of Fibonacci numbers to generate

// 	ch := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(1) // Increment the WaitGroup counter

// 	go fibonacci(n, ch, &wg) // Start the Fibonacci generation

// 	go func() {
// 		wg.Wait() // Wait for the Fibonacci generation to complete
// 		close(ch)
// 	}()

// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// }

package main

import (
	"fmt"
)

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
		quit <- false
	}(n)

	fibonacci(ch, quit)
}
