package main

import "fmt"

// 选择排序
func main() {
	a := []int{5, 3, 8, 0, 9, 1}
	fmt.Println()
	fmt.Printf("数组初始打印:%v\n", a)
	selectSort(a, len(a))
	fmt.Printf("数组排序结果:%v\n", a)
}

// 分为已排序区间和未排序区间，每次从未排序区间中找出最小值与待排序区间第一个元素交换
// 寻找最小值方法 设置初始状态最小值，从未排序区间中拿出元素进行比较，当小于最小值时则交换最小值下标及元素
func selectSort(a []int, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		minVal := a[i]
		for j := i + 1; j < n; j++ {
			if a[j] < minVal {
				minIndex = j
				minVal = a[j]
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}
