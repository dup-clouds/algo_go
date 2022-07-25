package _59_spiral_matrix_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, generateMatrix(3))
	ast.Equal([][]int{{1}}, generateMatrix(1))
}
