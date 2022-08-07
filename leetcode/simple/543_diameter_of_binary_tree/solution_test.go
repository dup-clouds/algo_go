package _543_diameter_of_binary_tree

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, 2, 3, 4, 5}
	root := structures.Nums2TreeNode(nums)
	ast.Equal(3, diameterOfBinaryTree(root))
}
