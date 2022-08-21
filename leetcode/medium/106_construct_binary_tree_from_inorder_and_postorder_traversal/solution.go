package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import "algo_go/structures"

type TreeNode = structures.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	treeMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		treeMap[v] = i
	}
	return build(treeMap, inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(treeMap map[int]int, inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	// 取后序遍历最后一个节点 即root节点
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}
	// 在中序遍历中寻找root节点
	index := treeMap[rootVal]
	// 计算左子树节点个数
	leftSize := index - inStart
	// 递归构建左子树 其中inorder的起点是inStart，终点是index-1即root节点左边的元素；其中postorder起始是postStart，终点是postStart+左节点个数-1即通过画图得出
	root.Left = build(treeMap, inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1)
	// 递归构建右子树
	root.Right = build(treeMap, inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return root
}
