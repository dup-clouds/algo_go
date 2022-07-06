package main

func main() {

}

type ListNodeCycleIi struct {
	Val  int
	Next *ListNodeCycleIi
}

/**
@link https://leetcode.cn/problems/linked-list-cycle-ii/
环形链表 II 查找循环链表的起点
1. 先找到环链表的相遇点位
2. 将慢指针指向head指针
3. 慢指针+快指针同时往前走，当相等时则为起点
设相遇点距离head距离为k，则慢指针走了k，快指针走了2k
设相遇点距离起点m 则head到起点为k-m，快指针到起点也为k-m
*/
func detectCycle(head *ListNodeCycleIi) *ListNodeCycleIi {
	slow := head
	fast := head
	for nil != fast && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
