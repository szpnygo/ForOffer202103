package main

import "fmt"

func main() {
	nums := []int{1, 0, 4, 0, 6}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := 0; j < len(nums)-1-i; j++ {
	// 		if nums[j] < nums[j+1] {
	// 			nums[j], nums[j+1] = nums[j+1], nums[j]
	// 		}
	// 	}
	// }
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
