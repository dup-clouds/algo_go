package main

import "fmt"

func main() {
	fmt.Println()
	a := []int{2, 7, 11, 15, 9}
	fmt.Println(twoSum4(a, 9))
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

// 暴力求解，从头开始，时间复杂度O(N)
func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
			if numbers[i]+numbers[j] > target {
				continue
			}
		}
	}
	return []int{0, 0}
}

// 遍历+二分查找，固定第一个数，使用二分查找寻找第二个数，时间复杂度O(Nlog(N))
func twoSum3(numbers []int, target int) []int {
	// 从0开始选择left元素
	for i := 0; i < len(numbers); i++ {
		// 求出right元素
		rightVal := target - numbers[i]
		// 二分查找低位
		low := i + 1
		// 二分查找高位
		high := len(numbers) - 1
		// 循环二分查找
		for low <= high {
			// 取中间值，需要加上基础变量low
			mid := (high-low)/2 + low
			if rightVal == numbers[mid] {
				return []int{i + 1, mid + 1}
			} else if rightVal > numbers[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return []int{0, 0}
}

// 双指针实现，取第一个元素和最后一个元素，依次向中间靠近，遍历完成数组即可找到对应下标，时间复杂度O(N)
func twoSum4(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		if numbers[low]+numbers[high] == target {
			return []int{low + 1, high + 1}
		} else if numbers[low]+numbers[high] < target {
			low++
		} else {
			high--
		}
	}
	return []int{0, 0}
}
