package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
}

/**
@link https://leetcode.cn/problems/remove-element/
27. 移除元素
*/
func removeElement(nums []int, val int) int {
	low := 0
	high := 0
	for high < len(nums) {
		if nums[high] != val {
			nums[low] = nums[high]
			low++
		}
		high++
	}
	fmt.Println(nums)
	return low
}
