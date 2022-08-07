package _144_binary_tree_preorder_traversal

import "algo_go/structures"

type TreeNode = structures.TreeNode

func preorderTraversal(root *TreeNode) (ans []int) {
	dfs(root, &ans)
	return
}
func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	dfs(root.Left, ans)
	dfs(root.Right, ans)
}
