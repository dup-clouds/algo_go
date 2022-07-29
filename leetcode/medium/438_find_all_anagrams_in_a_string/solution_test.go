package _438_find_all_anagrams_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{0, 6}, findAnagrams("cbaebabacd", "abc"))
	ast.Equal([]int{0, 1, 2}, findAnagrams("abab", "ab"))
}
