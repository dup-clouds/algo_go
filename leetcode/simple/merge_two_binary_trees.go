package main

import "fmt"

func main() {
	root1 := &TreeNode{1,
		&TreeNode{3,
			&TreeNode{5, nil,
				nil},
			nil},
		&TreeNode{2,
			nil,
			nil}}
	root2 := &TreeNode{2, &TreeNode{1, nil,
		&TreeNode{4, nil,
			nil}},
		&TreeNode{3, nil,
			&TreeNode{7, nil,
				nil}}}
	mergeRoot := mergeTrees(root1, root2)
	printRoot(mergeRoot)
}
func printRoot(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printRoot(root.Left)
	printRoot(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	newRoot := &TreeNode{root1.Val + root2.Val, nil, nil}
	newRoot.Left = mergeTrees(root1.Left, root2.Left)
	newRoot.Right = mergeTrees(root1.Right, root2.Right)
	return newRoot
}
