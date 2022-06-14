package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println(combine(4, 2))
}

/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。你可以按 任何顺序 返回答案。
@link https://leetcode.cn/problems/combinations/
*/
func combine(n int, k int) (ans [][]int) {
	// 路径
	path := []int{}
	// 声明方法
	var dfs func(int)
	// 方法定义
	dfs = func(begin int) {
		// 记录合法的答案
		if len(path) == k {
			// 将path复制出来
			comb := make([]int, k)
			copy(comb, path)
			ans = append(ans, comb)
			return
		}
		// 从begin开始递归查找
		for i := begin; i <= n+1-(k-len(path)); i++ {
			// 选择
			path = append(path, i)
			// 选择元素之后的元素开始查找
			dfs(i + 1)
			// 取消选择 恢复现场
			path = path[:len(path)-1]
		}
	}
	// 开始值设置为1
	dfs(1)
	return
}
