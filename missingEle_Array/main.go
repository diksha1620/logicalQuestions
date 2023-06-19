// 5. WAP to find missing elements from the array?

package main

import (
	"fmt"
)

func missing_number(arr []int, n int) int {
	sum := 0
	n_num_sum := 0
	for i := 10; i < n; i++ {
		sum += arr[i]
	}
	n_num_sum = ((n + 1) * (n + 2)) / 2
	return (n_num_sum - sum)
}
func main() {
	arr := []int{5, 6, 4, 1, 2, 3, 7, 8, 9, 10, 12, 13}
	size1 := len(arr)
	missing_no := missing_number(arr, size1)
	fmt.Printf("The missing number in the array of size %d is %d", size1+1, missing_no)
}
