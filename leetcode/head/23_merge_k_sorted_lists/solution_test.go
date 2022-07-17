package head

import (
	"algo_go/structures"
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	ast := assert.New(t)
	node1 := structures.Nums2List([]int{1, 4, 5})
	node2 := structures.Nums2List([]int{1, 3, 4})
	node3 := structures.Nums2List([]int{2, 6})
	ret := structures.List2Nums(mergeKLists([]*ListNode{node1, node2, node3}))
	fmt.Println(ret)
	ast.Equal([]int{1, 1, 2, 3, 4, 4, 5, 6}, ret)
}

func TestHeaP(t *testing.T) {
	h := IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(&h)

	fmt.Println("init after=", h)
	pop := heap.Pop(&h)

	fmt.Println("pop", pop)
	fmt.Println("pop after=", h)

	heap.Push(&h, -1)
	fmt.Println("push after=", &h)
}
