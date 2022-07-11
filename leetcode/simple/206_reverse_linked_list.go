package main

import "fmt"

func main() {
	head := &ListNodeRevers{1, &ListNodeRevers{2,
		&ListNodeRevers{3,
			&ListNodeRevers{4, &ListNodeRevers{5, nil}}}}}
	res := reverseList3(head)
	fmt.Println()
	printRevers(res)
}

type ListNodeRevers struct {
	Val  int
	Next *ListNodeRevers
}

func printRevers(res *ListNodeRevers) {
	for res == nil {
		return
	}
	fmt.Print(res.Val)
	printRevers(res.Next)
}

func reverseList(head *ListNodeRevers) *ListNodeRevers {
	var prev *ListNodeRevers
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseListLast(head *ListNodeRevers) *ListNodeRevers {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseListLast(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

/**
@link https://leetcode.cn/problems/reverse-linked-list/
反转链表
*/
func reverseList2(head *ListNodeRevers) *ListNodeRevers {
	if head == nil {
		return nil
	}
	ret := &ListNodeRevers{}
	count := 0
	for nil != head {
		if 0 == count {
			ret = nil
		}
		count++
		temp := &ListNodeRevers{head.Val, ret}
		ret = temp
		head = head.Next
	}
	return ret
}

// 反转链表递归实现
func reverseList3(head *ListNodeRevers) *ListNodeRevers {
	// 排除链表为空或只有一个节点情况
	if head == nil || head.Next == nil {
		return head
	}
	// 递归让链表反转 可看成 1 与 (2<-3<-4<-5)调整位置指向
	last := reverseList3(head.Next)
	// 设置当前节点的下下个节点指向自己，即反转相邻的两个节点
	head.Next.Next = head
	// 将当前节点的next节点指向空节点
	head.Next = nil
	printRevers(last)
	fmt.Println()
	return last
}
