package _889_construct_binary_tree_from_preorder_and_postorder_traversal

import "algo_go/structures"

type TreeNode = structures.TreeNode

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	treeMap := make(map[int]int, len(postorder))
	for i, v := range postorder {
		treeMap[v] = i
	}
	return build(treeMap, preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func build(treeMap map[int]int, preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	if preStart == preEnd {
		return &TreeNode{Val: preorder[preStart]}
	}
	// root节点值
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	// 第一个左子树节点值
	leftRootVal := preorder[preStart+1]
	index := treeMap[leftRootVal]
	// 计算root的左子节点的左树节点个数
	leftSize := index - postStart + 1
	// 递归构建左子节点 preorder起始位置为preStart+1即除去root节点， 截止位置preStart+左节点个数
	root.Left = build(treeMap, preorder, preStart+1, preStart+leftSize, postorder, postStart, index-1)
	// 递归构建右子树
	root.Right = build(treeMap, preorder, preStart+leftSize+1, preEnd, postorder, index+1, postEnd-1)
	return root
}
