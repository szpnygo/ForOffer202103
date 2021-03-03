package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(strconv.Itoa(reverse(1534236469)))
}

func reverse(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
