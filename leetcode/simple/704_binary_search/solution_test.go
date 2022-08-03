package _704_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	ast.Equal(-1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
