package main

import "fmt"

func main() {
	array := []int{1, 2, 3}
	fmt.Println(removeDuplicates(array))
}

/**
@link https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
26. 删除有序数组中的重复项
*/
func removeDuplicates(nums []int) int {
	// 已处理区间末尾
	low := 0
	// 快指针，没遇到重复元素，或非重复元素都+1，相当于未处理区间
	fast := 0
	for fast < len(nums) {
		if nums[low] != nums[fast] {
			low++
			nums[low] = nums[fast]
		}
		// 每循环一次加1
		fast++
	}
	return low + 1
}
