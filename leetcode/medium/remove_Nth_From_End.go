package main

import (
	"fmt"
)

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	fmt.Println()
	newHead := removeNthFromEnd(&head, 5)
	curr := newHead
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 删除链表的倒数第 N 个结点 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
* @link https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	curr := *head
	for curr.Next != nil {
		curr = *curr.Next
		length++
	}
	if n > length {
		return head
	}
	if n == length {
		head = head.Next
		return head
	}

	count := 0
	pre := head
	newCurr := head
	for newCurr.Next != nil {
		if count == length-n {
			break
		}
		pre = newCurr
		newCurr = newCurr.Next
		count++
	}
	pre.Next = pre.Next.Next
	return head
}
