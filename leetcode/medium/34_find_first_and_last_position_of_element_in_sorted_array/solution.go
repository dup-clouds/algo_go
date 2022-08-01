package _34_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) < 1 {
		return ans
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				ans[0] = mid
				break
			} else {
				high = mid - 1
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	low = 0
	high = len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			if mid+1 == len(nums) || nums[mid+1] != target {
				ans[1] = mid
				break
			} else {
				low = mid + 1
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
