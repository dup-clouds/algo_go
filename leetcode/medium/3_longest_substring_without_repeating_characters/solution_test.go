package _3_longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(3, lengthOfLongestSubstring("abcabcbb"))
	ast.Equal(1, lengthOfLongestSubstring("bbbbb"))
	ast.Equal(3, lengthOfLongestSubstring("pwwkew"))
}
