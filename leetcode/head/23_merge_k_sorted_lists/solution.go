package head

import (
	"algo_go/structures"
	"container/heap"
)

type ListNode = structures.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
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
	dummy := &ListNode{Val: -1}
	ret := dummy
	for len(h) > 0 {
		pVal := heap.Pop(&h)
		iAreaId, _ := pVal.(int)
		dummy.Next = &ListNode{Val: iAreaId}
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
