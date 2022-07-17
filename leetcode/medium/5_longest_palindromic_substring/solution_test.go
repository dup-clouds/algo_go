package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	ast := assert.New(t)
	str := "ababaa"
	ret := longestPalindrome(str)
	ast.Equal("ababa", ret)
}
func TestLongestPalindrome1(t *testing.T) {
	ast := assert.New(t)
	str := "a"
	ret := longestPalindrome(str)
	ast.Equal("a", ret)
}
func TestLongestPalindrome2(t *testing.T) {
	ast := assert.New(t)
	str := "ab"
	ret := longestPalindrome(str)
	ast.Equal("a", ret)
}
func TestLongestPalindrome3(t *testing.T) {
	ast := assert.New(t)
	str := ""
	ret := longestPalindrome(str)
	ast.Equal("", ret)
}
func TestLongestPalindrome4(t *testing.T) {
	ast := assert.New(t)
	str := "aa"
	ret := longestPalindrome(str)
	ast.Equal("aa", ret)
}
