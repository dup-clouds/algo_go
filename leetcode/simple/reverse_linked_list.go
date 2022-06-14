package main

import "fmt"

func main() {
	head := &ListNodeRevers{1, &ListNodeRevers{2, &ListNodeRevers{4, nil}}}
	res := reverseListLast(head)
	fmt.Println(res)
	//printRevers(res)
}

type ListNodeRevers struct {
	Val  int
	Next *ListNodeRevers
}

func printRevers(res *ListNodeRevers) {
	for res == nil {
		return
	}
	fmt.Println(res.Val)
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
