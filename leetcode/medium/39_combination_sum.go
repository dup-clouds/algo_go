package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{5, 3, 2}
	fmt.Println(combinationSum(candidates, 8))
}

/**
@link https://leetcode.cn/problems/combination-sum/
39. 组合总和
*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(candidates)
	backtrack(candidates, path, &res, target, 0, 0)
	return res
}

/**
candidates 选择列表
path 路径
res 结果集
target 目标值
sum 总和
start 开始选择位置，避免往前选择造成重复如：233 332 232 情况
*/
func backtrack(candidates []int, path []int, res *[][]int, target int, sum int, start int) {
	// 大于则直接返回
	if sum > target {
		return
	}
	// 采用加法法则 等于目标值时将path放入res中
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	// 遍历选择 从start开始，避免往前选择，造成重复数据
	for i := start; i < len(candidates); i++ {
		if candidates[i]+sum > target {
			break
		}
		// 做选择
		path = append(path, candidates[i])
		backtrack(candidates, path, res, target, sum+candidates[i], i)
		//撤销选择
		path = path[:len(path)-1]
	}
}
