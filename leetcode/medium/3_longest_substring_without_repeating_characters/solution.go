package _3_longest_substring_without_repeating_characters

import "math"

// 滑动窗口解决
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

// 暴力破解
func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maxVal := math.MinInt
	set := make(map[uint8]int)
	// 穷举所有可能
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if set[s[j]] < 1 {
				set[s[j]] = 1
			} else {
				break
			}
		}
		if len(set) > maxVal {
			maxVal = len(set)
		}
		set = make(map[uint8]int)
	}
	return maxVal
}

/**
核心思想：
max： 记录最大值，初始为0，每次滑动窗口时与r-i+1比较
i： 每次滑动窗口的起点
r： 作为一个右指针，从-1开始，不断叠加，不会重置，直到r+1=len(s)
每滑动一个窗口，则删除之前滑动过的元素，即当前滑动窗口初始下标-1
map存储的为每次滑动窗口不重复的字符，滑动时进行删除第一个元素
*/
func lengthOfLongestSubstring3(s string) int {
	// 散列表 存储当前枚举值
	m := make(map[byte]int)
	//
	r, maxCount := -1, 0
	// 滑动窗口起始位置
	for i := 0; i < len(s); i++ {
		// 删除窗口左边的元素，每次i+1则删除已使用过的元素
		if i != 0 {
			delete(m, s[i-1])
		}
		// 右指针小于s长度且对应s在map中无值时
		for r+1 < len(s) && m[s[r+1]] < 1 {
			m[s[r+1]]++
			r++
		}
		// 如果r已将整个串遍历完成并放入map中，则判断最大长度与当前i到len的长度，如果小于max，则无需后续设置maxCount步骤
		if r+1 == len(s) && r-i+1 < maxCount {
			break
		}
		if maxCount < r-i+1 {
			maxCount = r - i + 1
		}
	}
	return maxCount
}
