package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import "algo_go/structures"

type TreeNode = structures.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 构建中序遍历数组map，便于寻找root节点位置
	treeMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		treeMap[v] = i
	}
	return build(treeMap, preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
func build(treeMap map[int]int, preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	// 取到前序遍历preStart位置元素，即root节点
	rootVal := preorder[preStart]
	// 构建root节点
	root := &TreeNode{Val: rootVal}
	// 找到root节点对应中序遍历位置
	index := treeMap[rootVal]
	// 取到中序遍历中root节点到中序遍历开始出距离
	leftSize := index - inStart
	// 构建左子树
	root.Left = build(treeMap, preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
	// 构建右子树
	root.Right = build(treeMap, preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
	return root
}
