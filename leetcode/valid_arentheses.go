package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	s := "{[]}"
	fmt.Print(validParentheses(s))
}

// @link https://leetcode.cn/problems/valid-parentheses/
// 有效括号判断
func validParentheses(s string) bool {
	// 长度等于0或不是2的倍数则不符合
	if len(s)%2 == 1 {
		return false
	}
	// 定义括号规则，使用byte定义无需转字符串比较
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{'}
	// 定义栈
	var stack []byte
	for i := range s {
		// 存在有括号
		if pairs[s[i]] > 0 {
			// 判断栈是否为空 或者栈中最后一个元素是否与此时字符串中的括号是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 左括号加入栈中
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
