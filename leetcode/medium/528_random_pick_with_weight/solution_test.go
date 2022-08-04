package _528_random_pick_with_weight

import (
	"fmt"
	"testing"
)

func TestPickIndex(t *testing.T) {
	nums := []int{1, 3, 1, 1}
	s := Constructor(nums)
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
}
