package main

import "fmt"

func main() {
	head := &ListNodePartition{2, &ListNodePartition{1, nil}}
	newHead := partition(head, 2)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

type ListNodePartition struct {
	Val  int
	Next *ListNodePartition
}

/*
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。
@link https://leetcode.cn/problems/partition-list/
*/
func partition(head *ListNodePartition, x int) *ListNodePartition {
	if head == nil {
		return nil
	}
	lowHead := &ListNodePartition{-1, nil}
	tempLowHead := lowHead
	highHead := &ListNodePartition{-1, nil}
	tempHighHead := highHead
	for head != nil {
		if head.Val < x {
			lowHead.Next = &ListNodePartition{head.Val, nil}
			lowHead = lowHead.Next
		} else {
			highHead.Next = &ListNodePartition{head.Val, nil}
			highHead = highHead.Next
		}
		head = head.Next
	}
	lowHead.Next = tempHighHead.Next
	return tempLowHead.Next
}
