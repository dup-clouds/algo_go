package _104_maximum_depth_of_binary_tree

import "algo_go/structures"

type TreeNode = structures.TreeNode

// maxDepth 按层遍历记录深度-即广度优先遍历
func maxDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node != nil {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		queue = queue[size:]
		ans++
	}
	return ans
}

// maxDepth1 深度优先遍历
func maxDepth1(root *TreeNode) int {
	max := 0
	dfs(root, 0, &max)
	return max
}

func dfs(root *TreeNode, depth int, max *int) {
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

// maxDepth2 深度优先遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth2(root.Left)
	maxRight := maxDepth2(root.Right)
	res := maxVal(maxLeft, maxRight) + 1
	return res
}
