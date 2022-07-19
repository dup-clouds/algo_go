package _03_range_sum_query_immutable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumRange(t *testing.T) {
	ast := assert.New(t)
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)
	ast.Equal(numArray.SumRange(0, 2), 1)
	ast.Equal(numArray.SumRange(0, 0), -2)
	ast.Equal(numArray.SumRange(0, 5), -3)
	ast.Equal(numArray.SumRange(4, 5), 1)
}
