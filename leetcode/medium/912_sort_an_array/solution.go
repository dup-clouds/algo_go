package _912_sort_an_array

func sortArray(nums []int) []int {
	// 对0~n-1 位置进行归并排序
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, low, high int) {
	// 递归终止条件
	if low == high {
		return
	}
	// 取中间位置
	mid := low + (high-low)/2
	// 递归左侧
	sort(nums, low, mid)
	// 递归右侧
	sort(nums, mid+1, high)
	// 合并数组nums[low, mid] nums[mid+1, high]
	merge(nums, low, mid, high)
}

func merge(nums []int, low, mid, high int) {
	i := low
	j := mid + 1
	temp := make([]int, high-low+1)
	k := 0
	for i <= mid && j <= high {
		if nums[i] > nums[j] {
			temp[k] = nums[j]
			j++
			k++
		} else {
			temp[k] = nums[i]
			i++
			k++
		}
	}
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	for j <= high {
		temp[k] = nums[j]
		j++
		k++
	}
	for m := 0; m < len(temp); m++ {
		nums[m+low] = temp[m]
	}
}
