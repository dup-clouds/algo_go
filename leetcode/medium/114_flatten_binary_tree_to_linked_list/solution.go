package _114_flatten_binary_tree_to_linked_list

import (
	"algo_go/structures"
)

type TreeNode = structures.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dummyRoot := TreeNode{Val: -1}
	retRoot := &dummyRoot
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		retRoot.Right = &TreeNode{Val: node.Val, Left: nil}
		retRoot = retRoot.Right
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	*root = *dummyRoot.Right
}
