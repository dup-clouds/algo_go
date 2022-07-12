package main

import "fmt"

func main() {
	head := &ListNodeReverseBetween{1, &ListNodeReverseBetween{2,
		&ListNodeReverseBetween{3,
			&ListNodeReverseBetween{4, &ListNodeReverseBetween{5, nil}}}}}
	res := reverseBetween(head, 2, 4)
	fmt.Println()
	printRevers(res)
}
func printRevers(res *ListNodeReverseBetween) {
	for res == nil {
		return
	}
	fmt.Print(res.Val)
	printRevers(res.Next)
}

type ListNodeReverseBetween struct {
	Val  int
	Next *ListNodeReverseBetween
}

/**
@Link https://leetcode.cn/problems/reverse-linked-list-ii/
反转left到right之间的链表
*/
func reverseBetween(head *ListNodeReverseBetween, left int, right int) *ListNodeReverseBetween {
	if left == 1 {
		printRevers(head)
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	//printRevers(head)
	fmt.Println()
	return head
}

// 反转链表的前n个节点
func reverseN(head *ListNodeReverseBetween, n int) *ListNodeReverseBetween {
	// 保存链表第n+1个节点，便于后续反转后连接
	var successor *ListNodeReverseBetween
	// 定义递归函数
	var fun func(*ListNodeReverseBetween, int) *ListNodeReverseBetween

	// 递归函数实现
	fun = func(h *ListNodeReverseBetween, i int) *ListNodeReverseBetween {
		// 当只有一个节点时直接返回
		if i == 1 {
			successor = h.Next
			return h
		}
		// 递归下一个节点
		last := fun(h.Next, i-1)
		// 逐级反转
		h.Next.Next = h
		// 反转链表与未反转链表连接
		h.Next = successor
		// 返回第n个节点指针
		return last
	}

	return fun(head, n)
}
