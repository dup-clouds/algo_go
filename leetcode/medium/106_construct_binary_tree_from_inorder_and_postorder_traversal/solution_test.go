package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	ast := assert.New(t)
	inorder := []int{9, 3, 15, 20, 7}
	postOrder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postOrder)
	ast.Equal([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}, structures.Tree2Nums(root))
}
