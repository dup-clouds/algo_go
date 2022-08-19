package shu_zu_zhong_zhong_fu_de_shu_zi_lcof

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	ast := assert.New(t)
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	ast.Equal(2, findRepeatNumber(nums))
	ast.Equal(2, findRepeatNumber2(nums))
}
