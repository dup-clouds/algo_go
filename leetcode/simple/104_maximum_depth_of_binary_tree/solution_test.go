package _104_maximum_depth_of_binary_tree

import (
	"algo_go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	nums := []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}
	root := structures.Nums2TreeNode(nums)
	ast := assert.New(t)
	ast.Equal(3, maxDepth(root))
}
