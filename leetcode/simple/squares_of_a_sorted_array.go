package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(sortedSquaresLast([]int{-7, -3, 2, 3, 11}))
}

// 有序数组的平方
// @link https://leetcode.cn/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	n := len(nums)
	squares := make([]int, n)
	for i := 0; i < n; i++ {
		result := nums[i] * nums[i]
		j := i - 1
		for j >= 0 {
			if squares[j] > result {
				squares[j+1] = squares[j]
				j--
			} else {
				break
			}
		}
		squares[j+1] = result
	}
	return squares
}

// 优化
func sortedSquaresLast(nums []int) []int {
	n := len(nums)
	lastNegIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}
	resultArray := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 {
			resultArray = append(resultArray, nums[j]*nums[j])
			j++
		} else if j >= n {
			resultArray = append(resultArray, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			resultArray = append(resultArray, nums[i]*nums[i])
			i--
		} else {
			resultArray = append(resultArray, nums[j]*nums[j])
			j++
		}
	}
	return resultArray
}
