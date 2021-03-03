package main

import (
	"fmt"
)

func main() {
	fmt.Println("result：", longestPalindrome("abcba"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	result := string(s[1])
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		right := i
		left := i
		//找到重复字符串
		for r := 1; r < sLen-left; r++ {
			if s[left] == s[left+r] {
				right = left + r
				//len(string(s[left:right+1]))
				if right+1-left > len(result) {
					result = string(s[left : right+1])
				}
			} else {
				break
			}
		}
		//避免重复计算
		i = right
		//向两边扩展，一直到不能扩展为止
		for j := 1; j < sLen-i; j++ {
			left--
			right++
			if left >= 0 {
				if s[left] != s[right] {
					break
				} else {
					//len(string(s[left:right+1]))
					if right+1-left > len(result) {
						result = string(s[left : right+1])
					}
				}
			}
		}
	}
	return result
}
