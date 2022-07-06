package main

func main() {

}

type ListNodeIntersection struct {
	Val  int
	Next *ListNodeIntersection
}

/**
@link https://leetcode.cn/problems/intersection-of-two-linked-lists/
160. 相交链表
返回链表相交的节点
1. 将链表a拼接上链表b
2. 将链表b拼接上链表a
3. 遍历相等则相交与此节点
*/
func getIntersectionNode(headA, headB *ListNodeIntersection) *ListNodeIntersection {
	l1 := headA
	l2 := headB
	for l1 != l2 {

		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}

	}
	return l1
}
