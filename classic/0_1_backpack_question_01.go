package main

import "fmt"

func main() {
	maxSum := -1
	mem := [5][101]bool{}
	fmt.Println("回溯递归解法：")
	f1(0, 0, []int{20, 10, 5, 30, 22}, 5, 100, &maxSum, &mem)
	fmt.Println(maxSum)
	fmt.Println("动态规划解法-二维数组存储状态：")
	fmt.Println(f12([]int{20, 10, 5, 30, 22}, 5, 100))
	fmt.Println("动态规划解法-一维数组存储状态：")
	fmt.Println(f13([]int{2, 2, 4, 6, 3}, 5, 9))
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

/**
0-1背包问题-动态规划
使用一维数组存储
weight 物品重量集合
n 物品个数
w 背包总重量
*/
func f13(weight []int, n int, w int) int {
	// 状态存储，值为已记录的重量，每次装背包时以已有记录进行迭代求和
	states := make([]bool, w+1)
	// 重量为0已计算
	states[0] = true
	if w >= weight[0] {
		// 第一个物品重量已记录
		states[weight[0]] = true
	}
	// 选择剩余的n-1个物品进行迭代
	for i := 1; i < n; i++ {
		// 从后往前查找已记录的重量，防止重复使用，如果是从小到大则会出现：如j=0,weight[i]=2,4->true,6->true 当遍历到4时会导致重复计算一个物品
		for j := w - weight[i]; j >= 0; j-- {
			// 判断重量是否已记录过，记录过则可进行选择第i个物品相加
			if states[j] {
				states[j+weight[i]] = true
			}
		}
	}
	fmt.Printf("所有最大值可能性=%v", states)
	fmt.Println()
	// 从后往前遍历背包总重量，当存在已记录重量i时，则为最大背包可装值
	// 此处背包所有可能的最大值装法已包含
	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}
