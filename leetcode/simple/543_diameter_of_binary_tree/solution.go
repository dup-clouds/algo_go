package _543_diameter_of_binary_tree

import "algo_go/structures"

type TreeNode = structures.TreeNode

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	maxDepth(root, &ans)
	return
}

// maxDepth2 深度优先遍历
func maxDepth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left, ans)
	maxRight := maxDepth(root.Right, ans)
	*ans = maxVal(maxLeft+maxRight, *ans)
	res := maxVal(maxLeft, maxRight) + 1
	return res
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}
