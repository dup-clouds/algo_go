package structures

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNums2TreeNode(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := Nums2TreeNode(nums)
	newNums := Tree2Nums(root)
	ast.Equal(nums, newNums)
}

func TestGetTargetNode(t *testing.T) {
	nums := []int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}
	root := Nums2TreeNode(nums)

	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{

		{
			"找到 root.Right.Right",
			args{
				root:   root,
				target: 8,
			},
			root.Right.Right,
		},

		{
			"找到 root.Left.Left",
			args{
				root:   root,
				target: 6,
			},
			root.Left.Left,
		},

		//
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTargetNode(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTargetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree2Nums(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	root := Nums2TreeNode(nums)
	newNums := Tree2Nums(root)
	ast.Equal(nums, newNums)
}
func TestTreeNode_Equal(t *testing.T) {
	ast := assert.New(t)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	root := Nums2TreeNode(nums)
	root2 := Nums2TreeNode(nums2)
	ast.Equal(true, root.Equal(root2))
}
