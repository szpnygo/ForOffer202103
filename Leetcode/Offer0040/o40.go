package main

import "fmt"

func main() {
	fmt.Println(getLeastNumbers([]int{1, 3, 4, 5, 6}, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[:k]
}
