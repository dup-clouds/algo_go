package main

import "fmt"

func main() {
	head := ListNodeCycle{3, &ListNodeCycle{2,
		&ListNodeCycle{0, &ListNodeCycle{-4, nil}}}}
	fmt.Println(hasCycle(&head))
}

type ListNodeCycle struct {
	Val  int
	Next *ListNodeCycle
}

/**
@link https://leetcode.cn/problems/linked-list-cycle/
环形链表判断
*/
func hasCycle(head *ListNodeCycle) bool {
	mem := make(map[ListNodeCycle]bool)
	for head != nil {
		if mem[*head] {
			return true
		}
		mem[*head] = true
		head = head.Next
	}
	return false
}

// 升级版 判断链表中是否有环存在，快慢指针
func hasCycle1(head *ListNodeCycle) bool {
	// 慢指针
	slow := head
	// 快指针
	fast := head
	// 判断快指针及快指针的下一个节点
	for fast != nil && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
		// 如果存在换则最终将相遇
		if slow == fast {
			return true
		}
	}
	return false
}
