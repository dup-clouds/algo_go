package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	ast := assert.New(t)
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	ast.Equal([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}, structures.Tree2Nums(root))
}
