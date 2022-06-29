package main

import "fmt"

func main() {
	left := TreeNode5{9, nil, nil}
	right := TreeNode5{20, &TreeNode5{15, nil, nil}, &TreeNode5{7, nil, nil}}
	root := &TreeNode5{3, &left, &right}
	fmt.Println(pathSum(root, 12))
}

type TreeNode5 struct {
	Val   int
	Left  *TreeNode5
	Right *TreeNode5
}

/**
路径总和 II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
@link https://leetcode.cn/problems/path-sum-ii/
*/
func pathSum(root *TreeNode5, targetSum int) [][]int {
	// 存储最终结果
	var ans [][]int
	// 递归方法定义
	var pathFunc func(*TreeNode5, *[]int)
	pathFunc = func(node *TreeNode5, path *[]int) {
		// 节点为空时返回
		if nil == node {
			return
		}
		// 前序遍历 将节点值添加进路径中
		*path = append(*path, node.Val)
		// 判断是否为叶子节点--左节点+右节点为空
		// 计算路径和与目标值是否相等
		if node.Left == nil && node.Right == nil {
			sum := 0
			for _, v := range *path {
				sum += v
			}
			if sum == targetSum {
				// 临时变量，用于copy path值，避免使用引用而修改原值
				temp := make([]int, len(*path))
				copy(temp, *path)
				ans = append(ans, temp)
			}
		}
		// 递归遍历左节点
		pathFunc(node.Left, path)
		// 递归遍历右节点
		pathFunc(node.Right, path)
		// 出栈，删除最顶端元素
		*path = (*path)[:len(*path)-1]
	}
	var tempPath []int
	// 一次调用
	pathFunc(root, &tempPath)
	return ans
}
