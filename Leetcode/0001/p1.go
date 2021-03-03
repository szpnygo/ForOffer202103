package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	result := map[int]int{}
	for i, v := range nums {
		if val, ok := result[target-v]; ok {
			return []int{i, val}
		}
		result[v] = i
	}
	return []int{0, 0}
}
