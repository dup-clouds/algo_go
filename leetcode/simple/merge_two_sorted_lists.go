package main

import "fmt"

func main() {
	head1 := &ListNodeMerge{1, &ListNodeMerge{2, &ListNodeMerge{4, nil}}}
	head2 := &ListNodeMerge{1, &ListNodeMerge{3, &ListNodeMerge{4, nil}}}
	res := mergeTwoListsLast(head1, head2)
	fmt.Println()
	printRes(res)
}

type ListNodeMerge struct {
	Val  int
	Next *ListNodeMerge
}

func printRes(res *ListNodeMerge) {
	for res == nil {
		return
	}
	fmt.Println(res.Val)
	printRes(res.Next)
}

/**
合并两个有序链表
@link https://leetcode.cn/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(list1 *ListNodeMerge, list2 *ListNodeMerge) *ListNodeMerge {
	preHead := ListNodeMerge{-1, nil}
	prev := &preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return preHead.Next
}

/**
合并两个有序链表递归实现
1. l1为空，则返回l2
2. l2为空，则返回l1
3. l1.val<=l2.val 则l1为前一个元素，并递归设置l1的next节点
4. l2.val>l1.val 则递归设置l2的next节点
*/
func mergeTwoListsLast(list1 *ListNodeMerge, list2 *ListNodeMerge) *ListNodeMerge {
	if list1 == nil {
		return list2
	} else if nil == list2 {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsLast(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsLast(list1, list2.Next)
		return list2
	}
}
