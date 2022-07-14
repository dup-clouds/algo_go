package main

import "fmt"

func main() {
	head := &ListNodePalindrome{1,
		&ListNodePalindrome{2,
			&ListNodePalindrome{2,
				&ListNodePalindrome{1, nil}}}}
	res := isPalindrome2(head)
	fmt.Println(res)
}

type ListNodePalindrome struct {
	Val  int
	Next *ListNodePalindrome
}

/**
234 判断单链表是否是回文链表
@link https://leetcode.cn/problems/palindrome-linked-list/
放入数组中判断
*/
func isPalindrome(head *ListNodePalindrome) bool {
	var arrayVal []int
	for head != nil {
		arrayVal = append(arrayVal, head.Val)
		head = head.Next
	}
	n := len(arrayVal)
	a := arrayVal[:n/2]
	fmt.Println(5 / 2)
	fmt.Println(a)
	for i, v := range arrayVal[:n/2] {
		if v != arrayVal[n-i-1] {
			return false
		}
	}
	return true
}

// 通过快慢指针+反转链表实现
func isPalindrome2(head *ListNodePalindrome) bool {
	// 寻找需要反转的链表开头-即slow
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 奇数时需要往后走一个
	if fast != nil {
		slow = slow.Next
	}

	// 反转
	right := reverseHead(slow)

	// 判断val
	for right != nil {
		if right.Val != head.Val {
			return false
		}
		right = right.Next
		head = head.Next
	}

	return true
}
func reverseHead(head *ListNodePalindrome) *ListNodePalindrome {
	var pre *ListNodePalindrome
	for nil != head {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
