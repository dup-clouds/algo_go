package main

import (
	"fmt"
)

// @link https://leetcode.cn/problems/longest-common-prefix/
func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "dacecar", "dcar"}
	fmt.Println()
	fmt.Printf(longestCommonPrefix(strs))
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	// 取到数组第一个字符串
	str := strs[0]
	// 遍历初始字符串值
	for i := range str {
		// 遍历初始字符串外的其他字符串
		for j := 1; j < len(strs); j++ {
			// 当对应遍历字符串长度已遍历完成时 或当前字符串与初始字符串对应下标不同时，说明前缀字符串已达到最大值
			if i == len(strs[j]) || str[i] != strs[j][i] {
				return str[:i]
			}
		}
	}
	// 遍历字符串与初始字符串全部相等，则最大前缀字符串为初始字符串
	return str
}
