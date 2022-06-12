package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println(maxSubStringLen("cada"))
}

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
@link https://leetcode.cn/problems/longest-substring-without-repeating-characters/
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := make(map[string]int)
	for i := 0; i < len(s); i++ {
		if max[string(s[i])] < 1 {
			max[string(s[i])] = 1
		} else {
			break
		}
	}
	set := make(map[string]int)
	for i := 1; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if set[string(s[j])] < 1 {
				set[string(s[j])] = 1
			} else {
				break
			}
		}
		if len(set) > len(max) {
			max = set
		}
		set = make(map[string]int)
	}
	return len(max)
}

/**
核心思想：
max： 记录最大值，初始为0，每次滑动窗口时与r-i+1比较
i： 每次滑动窗口的步长
r： 作为一个右指针，从-1开始，不断叠加，不会重置，直达r+1=len(s)
每滑动一个窗口，则删除之前滑动过的元素，即当前滑动窗口初始下标-1
map存储的为每次滑动窗口不重复的字符，滑动时进行删除第一个元素
*/
func lengthOfLongestSubstringLast(s string) int {
	// 散列表 存储当前枚举值
	m := make(map[byte]int)
	//
	r, maxCount := -1, 0
	// 滑动窗口起始位置
	for i := 0; i < len(s); i++ {
		// 删除窗口左边的元素，每次i+1则从头开始删除一个
		if i != 0 {
			delete(m, s[i-1])
		}
		// 右指针小于s长度且对应s在map中无值时
		for r+1 < len(s) && m[s[r+1]] < 1 {
			m[s[r+1]]++
			r++
		}
		// 如果r已将整个串遍历完成并放入map中，则判断最大长度与当前i到len的长度，如果小于max，则无需后续滑动遍历
		if r+1 == len(s) && r-i+1 < maxCount {
			break
		}
		maxCount = max(maxCount, r-i+1)
	}
	return maxCount
}

/**
r从0开始版本实现
*/
func lengthOfLongestSubstringLastCopy(s string) int {
	// 散列表 存储当前枚举值
	m := make(map[byte]int)
	//
	r, maxCount := 0, 0
	// 滑动窗口起始位置
	for i := 0; i < len(s); i++ {
		// 删除窗口左边的元素，每次i+1则从头开始删除一个
		if i != 0 {
			delete(m, s[i-1])
		} else {
			m[s[0]]++
			r++
		}
		// 右指针小于s长度且对应s在map中无值时
		for r < len(s) && m[s[r]] < 1 {
			m[s[r]]++
			r++
		}
		// 如果r已将整个串遍历完成并放入map中，则判断最大长度与当前i到len的长度，如果小于max，则无需后续滑动遍历
		if r+1 == len(s) && r-i+1 < maxCount {
			break
		}
		maxCount = max(maxCount, r-i)
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
利用双子针完成无重复字符串的最大长度查找
1. 设置right，left初始值
2. 初始化窗口
3. 循环right->n 每次循环将maxVal与right-left进行取最大值
4. 存在重复元素时left向右移动一位
*/
func maxSubStringLen(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	right := 0
	left := 0
	maxVal := 0
	window := make(map[byte]int)
	for right < n {
		c := s[right]
		window[c]++
		right++
		for window[c] > 1 {
			window[s[left]]--
			left++
		}
		maxVal = max(maxVal, right-left)
	}
	return maxVal
}
