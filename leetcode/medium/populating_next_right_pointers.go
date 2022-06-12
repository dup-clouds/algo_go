package main

import "fmt"

func main() {
	a := &Node{1, &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil}, &Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil}, nil}
	connect(a)
	printRoot1(a)
}
func printRoot1(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	if root.Next != nil {
		fmt.Println(root.Next.Val)
	}
	printRoot1(root.Left)
	printRoot1(root.Right)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
@link https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
填充每个节点的下一个右侧节点指针
*/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connection(root)
	return root
}
func connection(root *Node) {
	if root.Left == nil {
		return
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connection(root.Left)
	connection(root.Right)
}
