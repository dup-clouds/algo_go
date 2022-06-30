package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange1([]int{1, 2, 5}, 100))
}

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
@link https://leetcode.cn/problems/coin-change/
*/
func coinChange(coins []int, amount int) int {
	// 存储每个阶段最小值，即下标为amount，值为最小硬币数组合
	dp := make([]int, amount+1)
	// 初始化一个非区间内的值
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	// 默认amount=0时需要0个硬币，base case
	dp[0] = 0
	// 从1开始迭代每个金额
	for i := 1; i < len(dp); i++ {
		// 所有硬币选择一遍
		for _, coin := range coins {
			// 金额减去当前硬币面值小于0则不可组成答案
			if i-coin < 0 {
				continue
			}
			// 取当前金额i最新硬币数组合
			if dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	// 判断是否为初始化填充值，是则说明无该金额组合
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 递归解决
func coinChange1(coins []int, amount int) int {
	mem := make([]int, amount+1)
	for i, _ := range mem {
		mem[i] = -888
	}
	return dp(coins, amount, &mem)
}

func dp(coins []int, amount int, mem *[]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if (*mem)[amount] != -888 {
		return (*mem)[amount]
	}
	minCount := math.MaxInt
	for _, coin := range coins {
		sub := dp(coins, amount-coin, mem)
		if sub == -1 {
			continue
		}
		if sub < minCount {
			minCount = sub + 1
		}
	}
	if minCount == math.MaxInt {
		(*mem)[amount] = -1
	}
	(*mem)[amount] = minCount
	return (*mem)[amount]
}
