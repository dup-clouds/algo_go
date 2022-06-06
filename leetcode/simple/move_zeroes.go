package main

import "fmt"

func main() {
	fmt.Println()
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}

// 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// @link https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	n := len(nums)
	count := 0
	j := 0
	for i := 0; i < n; i++ {
		value := nums[i]
		if value == 0 {
			count++
		} else {
			nums[j] = nums[i]
			j++
		}
	}
	for i := n - count; i < n; i++ {
		nums[i] = 0
	}
}
