package _76_minimum_window_substring

import "math"

func minWindow(s string, t string) string {
	// 待匹配字符串map，key-单个值，value-出现次数
	needMap := make(map[uint8]int)
	for i := range t {
		needMap[t[i]]++
	}
	// 窗口
	windowsMap := make(map[uint8]int)
	// 左边界
	left := 0
	// 右边界
	right := 0
	// 已验证值个数
	valid := 0
	// 匹配开始下标
	start := 0
	// 最终匹配长度
	length := math.MaxInt
	// 遍历整个字符串s
	for right < len(s) {
		// 取出有边界值
		curr := s[right]
		right++
		// 判断值是否在需要匹配的字符串中，若在则窗口增加
		if needMap[curr] > 0 {
			windowsMap[curr]++
			// 判断是否完全匹配目标字符串中的一类值
			if windowsMap[curr] == needMap[curr] {
				valid++
			}
		}
		// 判断是否存在最小匹配子串
		for valid == len(needMap) {
			// 缩小范围
			if right-left < length {
				start = left
				length = right - left
			}
			// 取出左边界字符
			leftVal := s[left]
			left++
			// 判断左边界字符是否在目标字符串中
			if needMap[leftVal] > 0 {
				if windowsMap[leftVal] == needMap[leftVal] {
					valid--
				}
				windowsMap[leftVal]--
			}
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
