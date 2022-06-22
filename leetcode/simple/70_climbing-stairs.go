package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
}

/**
@link https://leetcode.cn/problems/climbing-stairs/
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
