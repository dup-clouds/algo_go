package main

import "fmt"

func main() {
	left := TreeNode1{9, nil, nil}
	right := TreeNode1{20, &TreeNode1{15, nil, nil}, &TreeNode1{7, nil, nil}}
	root := &TreeNode1{3, &left, &right}
	fmt.Println(maxDepth1(root))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

/**
@link https://leetcode.cn/problems/maximum-depth-of-binary-tree/
求解二叉树的最大深度
*/
func maxDepth(root *TreeNode1) int {
	max := 0
	dfs(root, 0, &max)
	return max
}

func dfs(root *TreeNode1, depth int, max *int) {
	if nil == root {
		*max = maxVal(depth, *max)
		return
	}
	depth++
	dfs(root.Left, depth, max)
	dfs(root.Right, depth, max)
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth1(root *TreeNode1) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth1(root.Left)
	maxRight := maxDepth1(root.Right)
	res := maxVal(maxLeft, maxRight) + 1
	return res
}
