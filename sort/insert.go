package main

import "fmt"

// 插入排序
func main() {
	a := []int{5, 3, 8, 0, 9}
	fmt.Println()
	fmt.Printf("数组初始打印:%v\n", a)
	insertSort(a, len(a))
	fmt.Printf("数组排序结果:%v\n", a)
}

func insertSort(a []int, n int) {
	// 从未排序区间中取出一个元素，默认数组第一个元素设为有序
	for i := 1; i < n; i++ {
		// 取出待插入数据
		val := a[i]
		// 已排序区间最后一个元素下标
		j := i - 1
		// 从后往前遍历已排序区间
		for ; j >= 0; j-- {
			// 判断待排序元素与已排序区间最后一个元素大小关系,如待排序元素小于已排序区间最后一个元素,则将已排序区间最后一个元素后移一位
			if a[j] > val {
				a[j+1] = a[j]
			} else {
				// 不成立,说明已排序区间与待排序元素时从小到大关系
				break
			}
		}
		// 将待排序元素插入到适当位置
		a[j+1] = val
	}
}
