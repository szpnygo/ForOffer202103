package main

import "math"

func main() {
	reverse(-1234567)
}

func reverse(x int) int {
	n := 0
	for x != 0 {
		n = 10*n + x%10
		x = x / 10
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}
	return n
}
