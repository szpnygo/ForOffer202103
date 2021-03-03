package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(s string) int {
	maxSize := 0
	start := 0

	for i, v := range s {
		if index := strings.Index(s[start:i], string(v)); index != -1 {
			start = start + index + 1
		} else {
			if i+1-start > maxSize {
				maxSize = i + 1 - start
			}
		}

	}
	return maxSize
}
