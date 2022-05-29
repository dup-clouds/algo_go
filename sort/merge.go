package main

import "fmt"

// 归并排序
func main() {
	a := []int{666, 1, 0, 99, 99, 5, 3, 8, 0, 9}
	fmt.Println()
	fmt.Printf("数组初始打印:%v\n", a)
	mergeSort(a, 0, len(a)-1)
	fmt.Printf("数组排序结果:%v\n", a)
}

// 地推公式： mergeSort(p, r) = merge(mergeSort(p, q), mergeSort(q+1, r))
// 递归终止条件：p>=r
func mergeSort(a []int, p int, r int) {
	if p >= r {
		return
	}
	// 取中间节点位置
	q := p + ((r - p) >> 1)
	// 对a数组p~q进行排序
	mergeSort(a, p, q)
	// 对a数组q+1~r排序
	mergeSort(a, q+1, r)
	// 合并a[p~q] a[q+1~r]之前的数据到a中
	merge(a, p, q, r)
}

// 将a[p, q] a[p+1, r] 数据合并到a[p, r]中
func merge(a []int, p int, q int, r int) {
	// 游标i从左半部分开始
	i := p
	// 游标j从右半部分开始
	j := q + 1
	// 临时数组
	var tempArray = make([]int, 0)
	// 遍历左边和右边，其中一半遍历结束则停止遍历
	for i <= q && j <= r {
		if a[i] <= a[j] {
			tempArray = append(tempArray, a[i])
			i++
		} else {
			tempArray = append(tempArray, a[j])
			j++
		}
	}
	// 将未遍历完成的元素添加到临时数组
	start := i
	end := q
	if j <= r {
		start = j
		end = r
	}
	// 未遍历完成的元素重start到end [start, end]
	for i := start; i <= end; i++ {
		tempArray = append(tempArray, a[i])
	}
	// 将临时数组中的元素存储到a中，完成合并 其中a数组从p开始，到r-p结束
	for i = 0; i < len(tempArray); i++ {
		a[p+i] = tempArray[i]
	}
}
