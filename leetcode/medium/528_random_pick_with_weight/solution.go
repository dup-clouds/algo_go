package _528_random_pick_with_weight

import "math/rand"

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	n := len(w)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + w[i-1]
	}
	return Solution{preSum}
}

func (this *Solution) PickIndex() int {
	n := len(this.preSum)
	target := rand.Intn(this.preSum[n-1]) + 1
	return binarySearch(this.preSum, target) - 1
}

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
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
