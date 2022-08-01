package _34_find_first_and_last_position_of_element_in_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRange(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	ast.Equal([]int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	ast.Equal([]int{-1, -1}, searchRange([]int{}, 0))
}
