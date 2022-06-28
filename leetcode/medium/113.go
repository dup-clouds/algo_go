package main

import "fmt"

func main() {
	left := TreeNode5{9, nil, nil}
	right := TreeNode5{20, &TreeNode5{15, nil, nil}, &TreeNode5{7, nil, nil}}
	root := &TreeNode5{3, &left, &right}
	fmt.Println(pathSum(root, 12))
}

type TreeNode5 struct {
	Val   int
	Left  *TreeNode5
	Right *TreeNode5
}

func pathSum(root *TreeNode5, targetSum int) [][]int {
	var ans [][]int
	var pathFunc func(*TreeNode5, *[]int)
	pathFunc = func(node *TreeNode5, path *[]int) {
		if nil == node {
			return
		}
		*path = append(*path, node.Val)
		if node.Left == nil && node.Right == nil {
			sum := 0
			for _, v := range *path {
				sum += v
			}
			if sum == targetSum {
				temp := make([]int, len(*path))
				copy(temp, *path)
				ans = append(ans, temp)
			}
		}
		pathFunc(node.Left, path)
		pathFunc(node.Right, path)
		*path = (*path)[:len(*path)-1]
	}
	var tempPath []int
	pathFunc(root, &tempPath)
	return ans
}
