package main

import "fmt"

func main() {
	f2()
	fmt.Println("maxValue=", maxValue)
}

// 物品重量
var weight = []int{2, 2, 4, 6, 3}

// 物品价值
var value = []int{3, 4, 8, 9, 6}

// 物品个数
const n = 5

// 背包最大重量
const w = 9

// 最大价值
var maxValue = -1

/**
在背包总重量承载的情况下：
1. 尽可能装较多物品
2. 尽可能保持物品总价值最大--求出最大价值
*/
func f2() {
	// i-表示第n个物品，j-表示0-w的枚举，值存储价值
	states := [n][w + 1]int{}
	// 初始化成-1
	for i := 0; i < n; i++ {
		for j := 0; j < w+1; j++ {
			states[i][j] = -1
		}
	}
	// 选择不装第一个物品
	states[0][0] = 0
	if weight[0] <= w {
		// 选择装第一个物品
		states[0][weight[0]] = value[0]
	}
	// 选择第n个物品
	for i := 1; i < n; i++ {
		// 选择不装
		for j := 0; j < w; j++ {
			// 上一个物品+指定重量j存在时继承其价值
			if states[i-1][j] != -1 {
				states[i][j] = states[i-1][j]
			}
		}
		// 选择装入
		for j := 0; j <= w-weight[i]; j++ {
			// 上一个物品+指定重量j存在时 在此基础上加上重量和价值
			if states[i-1][j] != -1 {
				// 获取当前物品+前一个状态价值
				v := states[i-1][j] + value[i]
				// 新的价值大于已有数据时才更新，避免覆盖已有数据
				if v > states[i][j+weight[i]] {
					states[i][j+weight[i]] = v
				}
			}
		}
	}
	// 遍历最后一层取最大价值
	for m := w; m >= 0; m-- {
		if states[n-1][m] > maxValue {
			maxValue = states[n-1][m]
		}
	}
	// 打印输出二维数组
	fmt.Println()
	for i := 0; i < len(states); i++ {
		for j := 0; j < len(states[0]); j++ {
			fmt.Printf("%v  ", states[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
