package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println(lengthOfLongestSubstring("bbtablud"))
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
