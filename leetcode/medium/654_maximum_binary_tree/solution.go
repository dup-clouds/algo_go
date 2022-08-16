package _654_maximum_binary_tree

import (
	"algo_go/structures"
	"math"
)

type TreeNode = structures.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

// build nums-待构建数组 low-地位 high-高位
func build(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	// 获取low~high之间的最大值
	index := -1
	maxVal := math.MinInt
	for i := low; i <= high; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}
	// 构建root节点
	root := &TreeNode{Val: maxVal}
	// 递归构建左子树
	root.Left = build(nums, low, index-1)
	// 递归构建右子树
	root.Right = build(nums, index+1, high)
	return root
}
