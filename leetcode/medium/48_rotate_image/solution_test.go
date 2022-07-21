package _48_rotate_image

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	ast := assert.New(t)
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate2(nums)
	ast.Equal(nums, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})
}
