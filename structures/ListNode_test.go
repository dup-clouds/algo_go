package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_l2s(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{}, List2Nums(nil), "输入nil，没有返回[]int{}")

	one2three := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	ast.Equal([]int{1, 2, 3}, List2Nums(one2three), "没有成功地转换成[]int")

	limit := 100
	overLimitList := Nums2List(make([]int, limit+1))
	ast.Panics(func() { List2Nums(overLimitList) }, "转换深度超过 %d 限制的链条，没有 panic", limit)
}

func Test_s2l(t *testing.T) {
	ast := assert.New(t)
	ast.Nil(Nums2List([]int{}), "输入[]int{}，没有返回nil")

	ln := Nums2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	i := 1
	for ln != nil {
		ast.Equal(i, ln.Val, "对应的值不对")
		ln = ln.Next
		i++
	}
}

func Test_getNodeWith(t *testing.T) {
	ast := assert.New(t)
	ln := Nums2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	val := 10
	node := &ListNode{
		Val: val,
	}
	tail := ln
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	expected := node
	actual := ln.GetNodeWith(val)
	ast.Equal(expected, actual)
}

func Test_Nums2ListWithCycle(t *testing.T) {
	ast := assert.New(t)
	Nums := []int{1, 2, 3}
	l := Nums2ListWithCycle(Nums, -1)
	ast.Equal(Nums, List2Nums(l))

	l = Nums2ListWithCycle(Nums, 1)
	ast.Panics(func() { List2Nums(l) })
}
