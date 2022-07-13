package main

import (
	"fmt"
)

func main() {
	head := &ListNodeReverseBetween{1,
		&ListNodeReverseBetween{2,
			&ListNodeReverseBetween{3,
				&ListNodeReverseBetween{4,
					&ListNodeReverseBetween{5,
						&ListNodeReverseBetween{6,
							&ListNodeReverseBetween{7, nil}}}}}}}
	res := reverseBetween2(head, 3, 5)
	//res := reverse1(head)
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

// 使用迭代实现反转指定区间的链表 1 2 3 4 5 6 7
func reverseBetween2(head *ListNodeReverseBetween, left int, right int) *ListNodeReverseBetween {
	// 设置虚拟头结点 -1 1 2 3 4 5 6 7
	dummy := &ListNodeReverseBetween{-1, head}
	// 操作节点pre
	pre := dummy
	// 寻找left-1位置节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 寻找有边界节点 2 3 4 5 6 7->5 6 7
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	// 左边界节点赋值 2 3 4 5 6 7
	leftNode := pre.Next
	// 有边界节点的下一个节点赋值  6 7
	curr := rightNode.Next

	// 截取需要反转链表的头结点 -1 1 2 nil
	pre.Next = nil
	// 截取需要反转链表的尾结点 3 4 5 nil
	rightNode.Next = nil

	// 反转left-right节点 5 4 3
	reverse1(leftNode)

	// 拼接链表，进行还原 12 543 67
	pre.Next = rightNode
	leftNode.Next = curr

	// 返回已处理链表
	return dummy.Next
}

func reverse1(head *ListNodeReverseBetween) *ListNodeReverseBetween {
	// 定义前驱节点，默认为nil
	var prev *ListNodeReverseBetween
	// 遍历整个链表
	for head != nil {
		// 保存next指针
		next := head.Next
		// 对当前节点的next赋值给前驱节点
		head.Next = prev
		// 更新前驱节点
		prev = head
		// 更新当前节点为下一个节点
		head = next
	}
	return prev
}
