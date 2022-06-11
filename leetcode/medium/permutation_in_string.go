package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println("比较数组是否相等")
	a := [5]int{0, 0, 0, 1, 1}
	b := [5]int{0, 0, 1, 1, 1}
	fmt.Println(a == b)
}

/**
@link https://leetcode.cn/problems/permutation-in-string
字符串的排列
给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。
*/
func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	// n > m 说明s2中不包含s1任意排列的子串
	if n > m {
		return false
	}
	// 定义两个数组 大小限定26
	var cnt1, cnt2 [26]int
	// 遍历s1 将s1顺序放入数组中指定下标位置 将s2中0~len(s1)中元素放入数组中
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	// 判断两个数组是否相等 相等则条件成立
	if cnt1 == cnt2 {
		return true
	}
	// 到此则说明s2[0, len(s1)]中不包含s1排列字符串，则窗口需要往后移动一位，即s2最左边元素删除，s2最右边元素新增
	for i := n; i < m; i++ {
		// 增加窗口右边元素
		cnt2[s2[i]-'a']++
		// 去除窗口左边元素
		cnt2[s2[i-n]-'a']--
		// 判断是否相等
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
