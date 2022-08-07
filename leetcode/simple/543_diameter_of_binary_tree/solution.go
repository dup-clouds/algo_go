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
	// 动态root节点，计算ans
	*ans = maxVal(maxLeft+maxRight, *ans)
	// 返回最大深度，以左右子树深度最大值取得，考虑当前节点，故应加1
	res := maxVal(maxLeft, maxRight) + 1
	return res
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}
