package main

import "fmt"

func main() {
	a := []int{0, 0, 0, 0, 0, -1}
	fmt.Println()
	fmt.Println(reversePairs(a))
}

/**
@link https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数
*/
func reversePairs(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	ans := 0
	mergeSortCounting(&nums, 0, len(nums)-1, &ans)
	return ans
}
func mergeSortCounting(a *[]int, low int, high int, ans *int) {
	if low >= high {
		return
	}
	mid := (high + low) / 2
	mergeSortCounting(a, low, mid, ans)
	mergeSortCounting(a, mid+1, high, ans)
	merge(a, low, mid, high, ans)
	fmt.Println()
	fmt.Printf("a=%v,ans=%v", a, *ans)
	fmt.Println()
}

func merge(a *[]int, low int, mid int, high int, ans *int) {
	i := low
	j := mid + 1
	k := 0
	temp := make([]int, high-low+1)
	for i <= mid && j <= high {
		if (*a)[i] <= (*a)[j] {
			temp[k] = (*a)[i]
			k++
			i++
		} else {
			*ans = mid - i + 1 + *ans
			temp[k] = (*a)[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = (*a)[i]
		k++
		i++
	}
	for j <= high {
		temp[k] = (*a)[j]
		j++
		k++
	}
	for m, v := range temp {
		(*a)[low+m] = v
	}
}
