package main

import "fmt"

// @link https://leetcode.cn/problems/palindrome-number/
func main() {
	fmt.Println(palindrome(121))
}

// 判断一个int数字是否是回文字符串
// 取逆序整数进行比较 121 1*1+2*10+10*10
func palindrome(a int) bool {
	if a < 0 {
		return false
	}
	newVal := 0
	num := a
	for num != 0 {
		newVal = newVal*10 + num%10
		num = num / 10
	}
	return a == newVal
}
