package main

import (
	"fmt"
	"math"
)

func main() {
	left := TreeNode3{9, nil, nil}
	right := TreeNode3{20, &TreeNode3{15, nil, nil}, &TreeNode3{7, nil, nil}}
	root := &TreeNode3{3, &left, &right}
	fmt.Println(isBalanced(root))
}

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

// 判断二叉树是否为平衡二叉树
// @link https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode3) bool {
	flag := true
	balanced(root, 0, &flag)
	return flag
}

/*
总体思路：考虑后序遍历，及从下至上，最终到root节点
1. 获取左子树高度
2. 获取又子树高度
3. 判断左右子树相减绝对值是否大于1
4. 返回左右子树最大值+1
*/
func balanced(root *TreeNode3, depth int, flag *bool) int {
	if nil == root {
		return 0
	}
	leftDepth := balanced(root.Left, depth, flag)
	rightDepth := balanced(root.Right, depth, flag)
	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		*flag = false
		return 0
	}
	return max(leftDepth, rightDepth) + 1
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
