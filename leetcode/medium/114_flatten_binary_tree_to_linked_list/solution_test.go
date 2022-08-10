package _114_flatten_binary_tree_to_linked_list

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	ast := assert.New(t)
	treeNode := structures.Nums2TreeNode([]int{1, 2, 5, 3, 4, structures.NULL, 6})
	flatten(treeNode)
	ast.Equal([]int{1, structures.NULL, 2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6}, structures.Tree2Nums(treeNode))
}

func TestFlatten2(t *testing.T) {
	ast := assert.New(t)
	treeNode := structures.Nums2TreeNode([]int{0})
	flatten(treeNode)
	ast.Equal([]int{0}, structures.Tree2Nums(treeNode))
}
func TestFlatten3(t *testing.T) {
	ast := assert.New(t)
	treeNode := structures.Nums2TreeNode([]int{})
	flatten(treeNode)
	ast.Equal([]int{}, structures.Tree2Nums(treeNode))
}
