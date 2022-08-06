package _380_insert_delete_getrandom_o1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomizedSet_GetRandom(t *testing.T) {
	ast := assert.New(t)
	obj := Constructor()
	ast.Equal(true, obj.Insert(1))
	ast.Equal(true, obj.Insert(2))
	ast.Equal(true, obj.Insert(3))
	ast.Equal(true, obj.Insert(4))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}
func TestRandomizedSet_Insert(t *testing.T) {
	ast := assert.New(t)
	obj := Constructor()
	ast.Equal(true, obj.Insert(1))
	ast.Equal(true, obj.Insert(2))
	ast.Equal(true, obj.Insert(3))
	ast.Equal(true, obj.Insert(4))
	ast.Equal(false, obj.Remove(1))
}
func TestRandomizedSet_Remove(t *testing.T) {
	ast := assert.New(t)
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(4)
	ast.Equal(true, obj.Remove(3))
	ast.Equal(false, obj.Remove(5))
}
