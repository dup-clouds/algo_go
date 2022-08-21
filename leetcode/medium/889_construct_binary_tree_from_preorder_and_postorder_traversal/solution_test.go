package _889_construct_binary_tree_from_preorder_and_postorder_traversal

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructFromPrePost(t *testing.T) {
	ast := assert.New(t)
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postOrder := []int{4, 5, 2, 6, 7, 3, 1}
	root := constructFromPrePost(preorder, postOrder)
	ast.Equal([]int{1, 2, 3, 4, 5, 6, 7}, structures.Tree2Nums(root))
}
