package main

import "fmt"

func main() {
	fmt.Println()
	a := []int{-1, 0}
	fmt.Println(twoSum(a, -1))
}

// @link https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if numbers[i]+numbers[j] == target {
				if i < j {
					return []int{i + 1, j + 1}
				}
				return []int{j + 1, i + 1}
			}
		}
	}
	return []int{0, 0}
}
