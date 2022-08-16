package _654_maximum_binary_tree

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	ast := assert.New(t)
	nums := []int{3, 2, 1, 6, 0, 5}
	ast.Equal([]int{6, 3, 5, structures.NULL, 2, 0, structures.NULL, structures.NULL, 1}, structures.Tree2Nums(constructMaximumBinaryTree(nums)))
}
