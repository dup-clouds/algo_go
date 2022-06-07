package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	str := "Let's take LeetCode contest"
	fmt.Println()
	fmt.Println(reverseWordsLast(str))
}

// 反转字符串中的单词 III reverse-words-in-a-string-iii
// @link https://leetcode.cn/problems/reverse-words-in-a-string-iii/
func reverseWords(s string) string {
	sp := strings.Split(s, " ")
	resultStr := ""
	for i, s := range sp {
		array := strings.Split(s, "")
		low := 0
		high := len(s) - 1
		for low < high {
			array[low], array[high] = array[high], array[low]
			low++
			high--
		}
		if i == len(sp)-1 {
			resultStr += strings.Join(array, "")
		} else {
			resultStr += strings.Join(array, "") + " "
		}

	}
	return resultStr
}

// 优化反转
func reverseWordsLast(s string) string {
	n := len(s)
	ret := make([]string, n)
	for i := 0; i < n; {
		start := i
		for i < n && s[i] != ' ' {
			i++
		}
		for j := start; j < i; j++ {
			ret = append(ret, string(s[i-1-(j-start)]))
		}
		if i < n && s[i] == ' ' {
			i++
			ret = append(ret, " ")
		}
	}
	return strings.Join(ret, "")
}
