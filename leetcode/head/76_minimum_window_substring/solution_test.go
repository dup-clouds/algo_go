package _76_minimum_window_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	ast := assert.New(t)
	s := "ADOBECODEBANC"
	ta := "ABC"
	ast.Equal("BANC", minWindow(s, ta))
	s = "a"
	ta = "a"
	ast.Equal("a", minWindow(s, ta))
	s = "a"
	ta = "aa"
	ast.Equal("", minWindow(s, ta))
}
