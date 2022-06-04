package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(searchInsertLast([]int{1, 3, 5, 6}, 7))
}

// 考虑具有重复元素情况
func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		val := nums[mid]
		if val >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return len(nums)
}

// 无重复元素情况
// 搜索插入位置
// @link https://leetcode.cn/problems/search-insert-position/
func searchInsertLast(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		val := nums[mid]
		if val == target {
			return mid
		} else if val > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
