package main

import "fmt"

func main() {
	removeDuplicates([]int{1, 2, 3, 4, 4, 5, 5, 6, 7})
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			fmt.Println(i, j, nums)
		}
	}
	fmt.Println(nums)
	return i + 1
}

func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if total := prices[i] - prices[i-1]; total > 0 {
			sum += total
		}
	}
	return sum
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func containsDuplicate(nums []int) bool {
	cache := map[int]int{}
	for _, v := range nums {
		if _, ok := cache[v]; ok {
			return true
		}
		cache[v] = 1
	}
	return false
}

// func singleNumber(nums []int) int {
// 	n := 0
// 	cache := map[int]int{}
// 	for _, v := range nums {
// 		if _, ok := cache[v]; ok {
// 			n -= v
// 		} else {
// 			cache[v] = 1
// 			n += v
// 		}
// 	}
// 	return n
// }

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// func intersect(nums1 []int, nums2 []int) []int {

// }

func plusOne(digits []int) []int {
	n := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if n > 0 {
			digits[i] += n
			n--
		}
		if digits[i] >= 10 {
			digits[i] -= 10
			n++
		}
	}
	if n > 0 {
		digits = append([]int{n}, digits...)
	}
	return digits
}

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}

func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	for index, v := range nums {
		if i, ok := cache[v]; ok {
			return []int{i, index}
		}
		cache[target-v] = index
	}
	return []int{0, 0}
}
