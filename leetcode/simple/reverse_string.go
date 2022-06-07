package main

import "fmt"

func main() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println()
	fmt.Println(a)
	reverseString(a)
	fmt.Println(a)
}

// 反转字符串
// @link https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {
	lift := 0
	right := len(s) - 1
	for lift < right {
		s[lift], s[right] = s[right], s[lift]
		lift++
		right--
	}
}
