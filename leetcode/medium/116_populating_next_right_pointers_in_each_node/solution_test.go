package _116_populating_next_right_pointers_in_each_node

import "testing"

import (
	"fmt"
)

type question116 struct {
	para116
	ans116
}

// para 是参数
// one 代表第一个参数
type para116 struct {
	one *Node
}

// ans 是答案
// one 代表第一个答案
type ans116 struct {
	one *Node
}

func newQuestionNode() *Node {
	node7 := &Node{}
	node7.Val = 7

	node6 := &Node{}
	node6.Val = 6

	node5 := &Node{}
	node5.Val = 5

	node4 := &Node{}
	node4.Val = 4

	node3 := &Node{}
	node3.Val = 3

	node2 := &Node{}
	node2.Val = 2

	node1 := &Node{}
	node1.Val = 1

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7
	return node1
}

func newResultNode() *Node {
	node7 := &Node{}
	node7.Val = 7

	node6 := &Node{}
	node6.Val = 6

	node5 := &Node{}
	node5.Val = 5

	node4 := &Node{}
	node4.Val = 4

	node3 := &Node{}
	node3.Val = 3

	node2 := &Node{}
	node2.Val = 2

	node1 := &Node{}
	node1.Val = 1

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node1.Next = nil
	node2.Next = node3
	node3.Next = nil
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = nil

	return node1
}

func TestConnect1(t *testing.T) {

	qs := []question116{
		{
			para116{newQuestionNode()},
			ans116{newResultNode()},
		},
	}
	for _, q := range qs {
		_, p := q.ans116, q.para116
		fmt.Printf("【input】:%v      ", p.one)
		fmt.Printf("【output】:%v      \n", connect(p.one))
	}
	fmt.Printf("\n\n\n")
}

func TestConnect2(t *testing.T) {

	qs := []question116{
		{
			para116{newQuestionNode()},
			ans116{newResultNode()},
		},
	}
	for _, q := range qs {
		_, p := q.ans116, q.para116
		fmt.Printf("【input】:%v      ", p.one)
		fmt.Printf("【output】:%v      \n", connect2(p.one))
	}
	fmt.Printf("\n\n\n")
}
