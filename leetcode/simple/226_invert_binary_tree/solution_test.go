package _226_invert_binary_tree

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvertTree(t *testing.T) {
	ast := assert.New(t)
	nums := []int{4, 2, 7, 1, 3, 6, 9}
	ast.Equal([]int{4, 7, 2, 9, 6, 3, 1}, structures.Tree2Nums(invertTree(structures.Nums2TreeNode(nums))))
}

func TestInvertTree2(t *testing.T) {
	ast := assert.New(t)
	nums := []int{2, 1, 3}
	ast.Equal([]int{2, 3, 1}, structures.Tree2Nums(invertTree(structures.Nums2TreeNode(nums))))
}
