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

// 快速排序
func sortArray2(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}
func quicksort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	quicksort(nums, low, p-1)
	quicksort(nums, p+1, high)
}
func partition(nums []int, low, high int) int {
	partitionVal := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[j] < partitionVal {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

// sortArray3 冒泡排序
func sortArray3(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

// sortArray4 选择排序
func sortArray4(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		minVal := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < minVal {
				minVal = nums[j]
				minIndex = j
			}
		}
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}
	return nums
}

// sortArray5 插入排序
func sortArray5(nums []int) []int {
	// 待排序区间
	for i := 1; i < len(nums); i++ {
		// 待排序元素
		val := nums[i]
		// 从已排序区间末尾开始循环比较元素大小，并进行已排序区间元素后移
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 循环减一后需加回去
		nums[j+1] = val
	}
	return nums
}
