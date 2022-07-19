package _03_range_sum_query_immutable

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = nums[i] + sums[i]
	}
	return NumArray{sums}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.sums[right+1] - n.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
