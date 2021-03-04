package main

import "fmt"

func main() {
	nums := []int{1, 0, 4, 0, 6}
	fmt.Println(sortArray(nums))
}

func sortArray(nums []int) []int {
	// minIndex := 0
	// for i := 0; i < len(nums)-1; i++ {
	// 	minIndex = i
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[minIndex] > nums[j] {
	// 			minIndex = j
	// 		}
	// 	}

	// 	nums[i], nums[minIndex] = nums[minIndex], nums[i]
	// }
	// return nums

	for i := 1; i < len(nums); i++ {
		currentNum := nums[i]
		j := i - 1
		for j >= 0 && currentNum < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = currentNum
	}
	return nums
}
