package _3_remove_duplicates_from_sorted_list

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	ast := assert.New(t)
	head := structures.Nums2List([]int{1, 1, 2})
	ret := deleteDuplicates(head)
	ast.Equal(structures.List2Nums(ret), []int{1, 2})
}

func TestDeleteDuplicates1(t *testing.T) {
	ast := assert.New(t)
	head := structures.Nums2List([]int{1, 1, 2, 3, 3})
	ret := deleteDuplicates(head)
	ast.Equal(structures.List2Nums(ret), []int{1, 2, 3})
}

func TestDeleteDuplicates2(t *testing.T) {
	ast := assert.New(t)
	head := structures.Nums2List([]int{1, 1, 1, 1, 1})
	ret := deleteDuplicates(head)
	ast.Equal(structures.List2Nums(ret), []int{1})
}

func TestDeleteDuplicates3(t *testing.T) {
	ast := assert.New(t)
	head := structures.Nums2List([]int{1, 1, 1, 1, 2})
	ret := deleteDuplicatesLast(head)
	ast.Equal(structures.List2Nums(ret), []int{1, 2})
}
