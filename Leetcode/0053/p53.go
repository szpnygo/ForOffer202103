package main

func maxSubArray(nums []int) int {
	max := 0
	start := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		
	}
}
