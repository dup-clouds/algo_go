package main

import "fmt"

func main() {
	left := TreeNode7{9, nil, nil}
	right := TreeNode7{20, &TreeNode7{15, nil, nil}, &TreeNode7{7, nil, nil}}
	root := &TreeNode7{3, &left, &right}
	fmt.Println(hasPathSum3(root, 1111))
}

type TreeNode7 struct {
	Val   int
	Left  *TreeNode7
	Right *TreeNode7
}

/**
@link https://leetcode.cn/problems/path-sum/
路径总和
*/
func hasPathSum(root *TreeNode7, targetSum int) bool {
	flag := false
	dfsPath(root, &flag, targetSum)
	return flag
}

func dfsPath(root *TreeNode7, flag *bool, targetSum int) {
	if root == nil {
		return
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			*flag = true
			return
		}
	}
	dfsPath(root.Left, flag, targetSum)
	dfsPath(root.Right, flag, targetSum)
}

func hasPathSum1(root *TreeNode7, targetSum int) bool {
	if root == nil {
		return false
	}
	if nil == root.Left && nil == root.Right {
		return root.Val == targetSum
	}
	leftFlag := hasPathSum1(root.Left, targetSum-root.Val)
	if leftFlag {
		return true
	}
	return hasPathSum1(root.Right, targetSum-root.Val)
}

//广度优先遍历
func hasPathSum3(root *TreeNode7, targetSum int) bool {
	if nil == root {
		return false
	}
	var queNode []*TreeNode7
	var queVal []int
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) > 0 {
		node := queNode[0]
		queNode = queNode[1:]
		val := queVal[0]
		queVal = queVal[1:]
		if nil == node.Left && nil == node.Right {
			if targetSum == val {
				return true
			}
			continue
		}
		if nil != node.Left {
			queNode = append(queNode, node.Left)
			queVal = append(queVal, node.Left.Val+val)
		}
		if nil != node.Right {
			queNode = append(queNode, node.Right)
			queVal = append(queVal, node.Right.Val+val)
		}
	}
	return false
}
