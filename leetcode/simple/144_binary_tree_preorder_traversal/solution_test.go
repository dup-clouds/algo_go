package _144_binary_tree_preorder_traversal

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, structures.NULL, 2, 3}
	root := structures.Nums2TreeNode(nums)
	ast.Equal([]int{1, 2, 3}, preorderTraversal(root))
}
