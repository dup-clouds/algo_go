package main

import "fmt"

// 快速排序
func main() {
	a := []int{666, 1, 0, 99, 99, 5, 3, 8, 0, 9}
	fmt.Println()
	fmt.Printf("数组初始打印:%v\n", a)
	quicksort(a, 0, len(a)-1)
	fmt.Printf("数组排序结果:%v\n", a)
}

/**
快速排序思路
总体： 设排序数组下标p到r间的元素，选取一个分区点q,将小于q、大于q的元素分离,最后数组被分成三个部分(小于q、 q、大于q)
	  利用分治、递归处理思想，分别将p~q-1、q+1~r的元素进行分区，知道元素至少一个未知
递归公式：quicksort(p, r) = quicksort(p, q-1) + quicksort(q+1, r)
终止条件：p >= r
步骤：
1. 选择分区点，比如最后一个元素
2. 执行分区返回分区点
3. 执行p~q-1之前的分区
4. 执行q+1~r之间的分区
重点部分：分区函数
性质：原地排序算法、不稳定排序算法、时间复杂度为o(log n)
*/
func quicksort(a []int, p int, r int) {
	if p >= r {
		return
	}
	// 获取分区点q
	q := partition(a, p, r)
	quicksort(a, p, q-1)
	quicksort(a, q+1, r)
}

// 分区函数
func partition(a []int, p int, r int) int {
	// 分区点-取最后一个元素
	pivot := a[r]
	// 分区间隔点，a数组小于分区点的在i左边，大于分区点的在右边
	i := p
	// 从p到r-1遍历a数组
	for j := p; j < r; j++ {
		// 从未处理区间中拿出一个元素a[j]与分区点比较，小于分区点，则应该在已处理区间，即与已处理区间尾部元素i交换
		if a[j] < pivot {
			swap(a, i, j)
			// 已处理区间+1
			i++
		}
	}
	// 分区点与已处理区间交换
	swap(a, i, r)
	// 返回分区点下标
	return i
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
