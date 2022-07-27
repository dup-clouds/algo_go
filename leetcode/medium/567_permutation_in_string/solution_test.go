package _567_permutation_in_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(true, checkInclusion("ab", "eidbaooo"))
	ast.Equal(false, checkInclusion("ab", "eidboaoo"))
	ast.Equal(false, checkInclusion("hello", "ooolleoooleh"))
}
func TestCheckInclusion2(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(true, checkInclusion2("ab", "eidbaooo"))
	ast.Equal(false, checkInclusion2("ab", "eidboaoo"))
	ast.Equal(false, checkInclusion2("hello", "ooolleoooleh"))
}
