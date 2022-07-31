package _3_longest_substring_without_repeating_characters

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	// 存储已经遍历的字符
	windowsMap := make(map[uint8]int)
	// 窗口左边界
	left := 0
	// 窗口右边界
	right := 0
	// 返回结果
	maxAns := math.MinInt8
	// 遍历s
	for right < len(s) {
		// 取出当前遍历的字符
		curr := s[right]
		// 遍历游标+1
		right++
		// 窗口增加
		windowsMap[curr]++
		// 判断当前字符是否在窗口中，若在则缩小窗口，此处循环处理，避免多个重复值情况
		for windowsMap[curr] > 1 {
			leftVal := s[left]
			left++
			windowsMap[leftVal]--
		}
		// 判断是否更新结果
		if right-left > maxAns {
			maxAns = right - left
		}
	}
	return maxAns
}
