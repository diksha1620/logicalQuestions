// anonymous function and goroutines
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			fmt.Printf("Goroutine: %d\n", i)
// 			time.Sleep(time.Millisecond * 400)
// 		}
// 	}()

// 	time.Sleep(time.Second)
// 	fmt.Println("Main function")
// }

// anonymous function and channels

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // Send data to the channel
			time.Sleep(time.Millisecond * 500)
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}

	fmt.Println("Main function")
}
