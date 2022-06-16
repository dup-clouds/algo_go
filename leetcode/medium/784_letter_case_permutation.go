package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(letterCasePermutation("a1b2"))
	//fmt.Println(strings.Split("a1b2", "")[0])
}

/**
@link https://leetcode.cn/problems/letter-case-permutation/
给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
*/
func letterCasePermutation(s string) []string {
	var ansPer []string
	var dfs func([]byte, int)
	dfs = func(str []byte, index int) {
		temp := string(str)
		ansPer = append(ansPer, temp)
		for i := index; i < len(str); i++ {
			if str[i] > 47 && str[i] < 58 {
				continue
			}
			if str[i] > 64 && str[i] < 91 {
				str[i] += 32
				dfs(str, i+1)
				str[i] -= 32
			} else {
				str[i] -= 32
				dfs(str, i+1)
				str[i] += 32
			}

		}
	}
	dfs([]byte(s), 0)
	return ansPer
}
