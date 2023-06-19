// 1. WAP to find Second largest element in an array : arr[] = [12, 35, 1, 10, 34, 1, 35],
// without sorting, without using any built-in methods and without deleting duplicate elements.
// What will be the time complexity?

package main

import "fmt"

func main() {
	var large_1 int = 0
	var large_2 int = 0
	var array = [...]int{12, 35, 1, 10, 34, 1, 35}
	large_1 = array[0]
	for i := 0; i < len(array); i++ {
		if large_1 < array[i] {
			large_2 = large_1
			large_1 = array[i]
		} else if large_2 < array[i] {
			large_2 = array[i]
		}
	}
	fmt.Println("Second largest element is: ", large_2)
}

// Time Complexity = O(nlogn)
