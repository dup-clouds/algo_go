package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println()
	fmt.Print(binarySearch(nums, 9))
}

// 二分查找
// @link https://leetcode.cn/problems/first-bad-version/
func binarySearch(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
