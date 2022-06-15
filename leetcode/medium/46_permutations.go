package main

import "fmt"

func main() {
	fmt.Println()
	a := []int{0, 1}
	fmt.Println(permute(a))
}
func permute(nums []int) [][]int {
	dfsPermute(nums, []int{})
	return ans
}

var ans [][]int

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
@link https://leetcode.cn/problems/permutations/
nums 数组-动态变化，选择之后将删除自己，回溯时恢复
dfsPath 选择路径，回溯时移除路径最后一个元素，重新选择时为空
*/
func dfsPermute(nums []int, dfsPath []int) {
	// nums全被选择后则说明已经为完整路径，即可添加到答案中
	if len(nums) == 0 {
		tempPath := make([]int, len(dfsPath))
		copy(tempPath, dfsPath)
		ans = append(ans, tempPath)
	}
	// 遍历选择列表
	for i := 0; i < len(nums); i++ {
		// 选择当前值
		cur := nums[i]
		// 添加到路径中
		dfsPath = append(dfsPath, cur)
		// 选择列表中移除已选元素，避免重复选择
		nums = append(nums[:i], nums[i+1:]...)
		// 递归选择
		dfsPermute(nums, dfsPath)
		// 将选择重新添加到列表中
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		// 回溯恢复现场
		dfsPath = dfsPath[:len(dfsPath)-1]
	}
}
