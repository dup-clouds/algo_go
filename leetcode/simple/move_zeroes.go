package main

import "fmt"

func main() {
	fmt.Println()
	a := []int{0, 1, 0, 3, 12}
	moveZeroesLast(a)
	fmt.Println(a)
}

// 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// @link https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		value := nums[i]
		if value != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < n; i++ {
		nums[i] = 0
	}
}

// 优化-一次编辑即可，参考快排已排序区间和未排序区间
func moveZeroesLast(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			if i > j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
	}
}
