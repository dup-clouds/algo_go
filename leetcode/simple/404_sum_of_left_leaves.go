package main

import "fmt"

func main() {

	left := TreeNode6{9, nil, nil}
	right := TreeNode6{20, &TreeNode6{15, nil, nil}, &TreeNode6{7, nil, nil}}
	root := &TreeNode6{3, &left, &right}
	fmt.Println(sumOfLeftLeaves(root))
}

type TreeNode6 struct {
	Val   int
	Left  *TreeNode6
	Right *TreeNode6
}

/**
二叉树左叶子之和
@link https://leetcode.cn/problems/sum-of-left-leaves/
*/
func sumOfLeftLeaves(root *TreeNode6) int {
	ret := 0
	var dfs func(*TreeNode6, *int)
	dfs = func(node *TreeNode6, i *int) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			*i = *i + node.Left.Val
		}
		dfs(node.Left, i)
		dfs(node.Right, i)
	}
	dfs(root, &ret)
	return ret
}
