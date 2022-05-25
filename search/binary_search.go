package main

import "fmt"

func main() {
	//a := []int{3, 4, 6, 7, 10}
	array := []int{3, 5, 6, 8, 9, 10}
	// 非递归实现二分查找
	index := simpleBinarySearch(array, len(array), 5)
	// 递归实现二分查找
	indexRecursion := simpleBinarySearchByRecursion(array, 0, len(array)-1, 8)
	// 查找第一个等于给定值的元素位置
	arrayFirstEqual := []int{3, 5, 6, 8, 9, 9, 10}
	indexFirstEqual := bSearchFirstEqual(arrayFirstEqual, len(arrayFirstEqual), 9)
	// 查找最后一个等于给定值的元素位置
	arrayEndEqual := []int{3, 5, 6, 8, 9, 9, 10}
	indexEndEqual := bSearchEndEqual(arrayEndEqual, len(arrayEndEqual), 9)
	// 查找第一个大于等于给定值的元素位置 3 4 6 7 10 给定值：5 位置：2
	arrayFirstGtEqual := []int{3, 4, 6, 7, 10}
	indexGtEqual := bSearchFirstGtEqual(arrayFirstGtEqual, len(arrayFirstGtEqual), 5)
	// 查找最后一个小于等于给定值的元素位置 3 5 6 8 9 10 给定值：7 位置：2
	arrayEndLtEqual := []int{3, 5, 6, 8, 9, 10}
	indexEndLtEqual := bSearchEndLtEqual(arrayEndLtEqual, len(arrayEndLtEqual), 7)
	fmt.Println()
	fmt.Printf("非递归实现二分查找-数组=%v;预期结果=%v;实际结果=%v\n", array, 1, index)
	fmt.Printf("递归实现二分查找-数组=%v;预期结果=%v;实际结果=%v\n", array, 3, indexRecursion)
	fmt.Printf("查找第一个等于给定值的元素位置-数组=%v;预期结果=%v;实际结果=%v\n", arrayFirstEqual, 4, indexFirstEqual)
	fmt.Printf("查找最后一个等于给定值的元素位置-数组=%v;预期结果=%v;实际结果=%v\n", arrayEndLtEqual, 5, indexEndEqual)
	fmt.Printf("查找第一个大于等于给定值的元素位置-数组=%v;预期结果=%v;实际结果=%v\n", arrayFirstGtEqual, 2, indexGtEqual)
	fmt.Printf("查找最后一个小于等于给定值的元素位置-数组=%v;预期结果=%v;实际结果=%v\n", arrayEndLtEqual, 2, indexEndLtEqual)
}

// 简单for循环二分查找
func simpleBinarySearch(a []int, len int, value int) int {
	if len < 1 {
		return -1
	}
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == value {
			return mid
		} else if a[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 递归简单for循环二分查找
func simpleBinarySearchByRecursion(a []int, low int, high int, value int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if a[mid] == value {
		return mid
	} else if a[mid] > value {
		return simpleBinarySearchByRecursion(a, low, mid-1, value)
	} else {
		return simpleBinarySearchByRecursion(a, mid+1, high, value)
	}
}

// 二分查找第一个等于给定值的元素
func bSearchFirstEqual(a []int, len int, value int) int {
	if len < 1 {
		return -1
	}
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == value {
			if mid == 0 || a[mid-1] != value {
				return mid
			} else {
				high = mid - 1
			}
		} else if a[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个等于给定值的元素
func bSearchEndEqual(a []int, len int, value int) int {
	if len < 1 {
		return -1
	}
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] == value {
			if mid == len-1 || a[mid+1] != value {
				return mid
			} else {
				low = mid + 1
			}
		} else if a[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 二分查找第一个大于等于给定值的元素
func bSearchFirstGtEqual(a []int, len int, value int) int {
	if len < 1 {
		return -1
	}
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] >= value {
			if mid == 0 || a[mid-1] < value {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于给定值的元素
func bSearchEndLtEqual(a []int, len int, value int) int {
	if len < 1 {
		return -1
	}
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if a[mid] <= value {
			if mid == len-1 || a[mid+1] > value {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
