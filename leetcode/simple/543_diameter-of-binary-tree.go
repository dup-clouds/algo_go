package main

import "fmt"

func main() {
	right := TreeNode2{3, nil, nil}
	left := TreeNode2{2, &TreeNode2{4, nil, nil}, &TreeNode2{5, nil, nil}}
	root := &TreeNode2{1, &left, &right}
	fmt.Println(diameterOfBinaryTree(root))
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

/**
二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
@link https://leetcode.cn/problems/diameter-of-binary-tree/
*/
func diameterOfBinaryTree(root *TreeNode2) int {
	max := 0
	maxDepthDiameter(root, &max)
	return max
}
func maxDepthDiameter(root *TreeNode2, max *int) int {
	if root == nil {
		return 0
	}
	left := maxDepthDiameter(root.Left, max)
	right := maxDepthDiameter(root.Right, max)
	*max = maxDiameter(*max, left+right)
	return 1 + maxDiameter(left, right)

}
func maxDiameter(x, y int) int {
	if x > y {
		return x
	}
	return y
}
