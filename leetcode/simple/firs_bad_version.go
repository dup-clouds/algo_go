package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println(firstBadVersion(5))
}

/**
* Forward declaration of isBadVersion API.
* @param   version   your guess about first bad version
* @return 	 	      true if current version is bad
*			          false if current version is good
* func isBadVersion(version int) bool;
 */

// 找出第一个错误版本
// @link https://leetcode.cn/problems/first-bad-version/
func firstBadVersion(n int) int {
	if n < 1 {
		return -1
	}
	low := 1
	high := n
	for low <= high {
		mid := low + ((high - low) >> 1)
		bad := isBadVersion(mid)
		if bad {
			if mid == 1 || !isBadVersion(mid-1) {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	if version == 4 || version == 3 {
		return true
	}
	return false
}
