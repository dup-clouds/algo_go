package main

import (
	"container/heap"
	"fmt"
)

func main() {
	l1 := &ListNodeMergeK{1, &ListNodeMergeK{4, &ListNodeMergeK{5, nil}}}
	l2 := &ListNodeMergeK{1, &ListNodeMergeK{3, &ListNodeMergeK{4, nil}}}
	l3 := &ListNodeMergeK{2, &ListNodeMergeK{6, nil}}
	array := []*ListNodeMergeK{l1, l2, l3}
	ret := mergeKLists(array)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
	//h := IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	//heap.Init(&h)
	//fmt.Println(h)
	//pop := heap.Pop(&h)
	//fmt.Println(pop)
	//fmt.Println(h)
	//heap.Push(&h, 6)
}

type ListNodeMergeK struct {
	Val  int
	Next *ListNodeMergeK
}

/**
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
使用golang自带的堆实现最小堆逻辑
@link https://leetcode.cn/problems/merge-k-sorted-lists/
*/
func mergeKLists(lists []*ListNodeMergeK) *ListNodeMergeK {
	if len(lists) < 1 {
		return nil
	}
	h := IntHeap{}
	for _, v := range lists {
		head := v
		for head != nil {
			h = append(h, head.Val)
			head = head.Next
		}
	}
	heap.Init(&h)
	dummy := &ListNodeMergeK{-1, nil}
	ret := dummy
	for len(h) > 0 {
		pVal := heap.Pop(&h)
		iAreaId, _ := pVal.(int)
		dummy.Next = &ListNodeMergeK{iAreaId, nil}
		dummy = dummy.Next
	}
	return ret.Next
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
