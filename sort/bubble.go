package main

import "fmt"

// 冒泡排序
func main() {
	a := []int{3, 1, 9, 5, 0, 1, 2, 99}
	fmt.Println()
	fmt.Printf("排序前数组为=%v\n", a)
	bubble(a, len(a))
	fmt.Printf("排序后数组为=%v", a)
}

func bubble(a []int, n int) {
	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				swapBubble(a, j, j+1)
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

func swapBubble(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
