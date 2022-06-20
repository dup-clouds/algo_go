package main

import "fmt"

func main() {
	maxSum := -1
	mem := [5][101]bool{}
	fmt.Println("回溯递归解法：")
	f1(0, 0, []int{20, 10, 5, 30, 49}, 5, 100, &maxSum, &mem)
	fmt.Println(maxSum)
	fmt.Println("动态规划解法：")
	fmt.Println(f12([]int{20, 10, 5, 30, 49}, 5, 100))
}

/**
备忘录模式，使用二维数组记录对应物品-重量是否被计算过
*/
func f1(index int, cw int, items []int, n int, w int, maxSum *int, mem *[5][101]bool) {
	if cw == w || index == n {
		if cw > *maxSum {
			*maxSum = cw
		}
		return
	}
	if mem[index][cw] {
		return
	}
	mem[index][cw] = true
	f1(index+1, cw, items, n, w, maxSum, mem)
	if items[index]+cw <= w {
		f1(index+1, cw+items[index], items, n, w, maxSum, mem)
	}
}

/**
动态规划解法：
weight数组，存储每个物品的质量
n 物品个数
w 背包总质量
*/
func f12(weight []int, n int, w int) int {
	// i-第i个物品，j-加上上一次选择的重量和
	var states = [5][101]bool{}
	// 第一个物品选择不装
	states[0][0] = true
	if weight[0] <= w {
		// 第一个物品选择装
		states[0][weight[0]] = true
	}
	// 剩下的n-1个物品装入
	for i := 1; i < n; i++ {
		// 第i个物品不装
		for j := 0; j < w; j++ {
			// 当且仅当前一个物品的j重量存在时继承已有物品总和j
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		// 第i个物品选择装，且装的极限为w-weight[i]，即当前装入的重量+已有重量不能超过总重量，k为已有重量的枚举
		for k := 0; k <= w-weight[i]; k++ {
			// 上一个物品存在k重量时装入，不存在则无需装入
			if states[i-1][k] {
				states[i][weight[i]+k] = true
			}
		}
	}
	// 从总重量往前遍历，找出最后一层物品最接近w时即为最大值
	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}
	return 0
}
