package _3_remove_duplicates_from_sorted_list

import (
	"algo_go/structures"
)

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: -1, Next: head}
	slow := head
	fast := head.Next
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 特殊处理尾结点相同情况
	if slow != nil && slow.Next != nil && (slow.Val == slow.Next.Val) {
		slow.Next = nil
	}
	return dummy.Next
}

func deleteDuplicatesLast(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next != nil {
		// 判断当前值与下一个值是否相等
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
