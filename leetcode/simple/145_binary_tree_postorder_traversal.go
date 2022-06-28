package main

import (
	"fmt"
)

func main() {
	left := TreeNode4{9, nil, nil}
	right := TreeNode4{20, &TreeNode4{15, nil, nil}, &TreeNode4{7, nil, nil}}
	root := &TreeNode4{3, &left, &right}
	fmt.Println(postorderTraversal2(root))
}

type TreeNode4 struct {
	Val   int
	Left  *TreeNode4
	Right *TreeNode4
}

// 递归法
func postorderTraversal(root *TreeNode4) (ans []int) {
	var postFunc func(node *TreeNode4)
	postFunc = func(node *TreeNode4) {
		if node == nil {
			return
		}
		postFunc(node.Left)
		postFunc(node.Right)
		ans = append(ans, node.Val)
	}
	postFunc(root)
	return
}

// 非递归实现方法
// 使用栈记录 根->右->左
// 反转结果为 左->右->根
func postorderTraversal2(root *TreeNode4) (ans []int) {
	if nil == root {
		return nil
	}
	var res []int
	var stack []*TreeNode4
	stack = append(stack, root)
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		res = append(res, curr.Val)
		stack = stack[:len(stack)-1]
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
	low := 0
	high := len(res) - 1
	for low <= high {
		res[low], res[high] = res[high], res[low]
		high--
		low++
	}
	return res
}
