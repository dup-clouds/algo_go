package main

import "fmt"

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	fmt.Println()
	//curr := &head
	//for curr != nil {
	//	fmt.Println(curr.Val)
	//	curr = curr.Next
	//}
	fmt.Println(middleNode(&head).Val)
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 链表的中间结点 两个指针，第一个指针每次走一步，第二个指针每次走两步
 * @link https://leetcode.cn/problems/middle-of-the-linked-list/
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 获取下一个节点
	nextNode := head.Next
	if nextNode == nil {
		return head
	}
	// 获取下下个节点
	nextNextNode := head.Next.Next
	// 判断下下个节点是否为空且下下下个节点不为空时进行移动操作
	for nextNextNode != nil && nextNextNode.Next != nil {
		nextNode = nextNode.Next
		nextNextNode = nextNextNode.Next.Next
	}
	return nextNode
}
