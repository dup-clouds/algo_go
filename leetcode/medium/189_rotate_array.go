package main

import "fmt"

func main() {
	fmt.Println()
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotateLast(a, 3)
	fmt.Println(a)
}

// 数组旋转 向右移动第k位
// @link https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	n := len(nums)
	newArray := make([]int, n)
	for i := 0; i < n; i++ {
		j := (i + k) % n
		newArray[j] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = newArray[i]
	}
}

// 利用数组翻转特性
// 数组整体翻转一次
// 数组从0~k-1翻转一次
// 数组从k到n-1翻转一次
// 总的时间复杂度为O(2n)=O(n)
func rotateLast(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
func reverse(a []int, start, end int) {
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}
