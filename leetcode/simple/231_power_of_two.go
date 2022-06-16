package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo(5))
}

/**
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。
@link https://leetcode.cn/problems/power-of-two/
*/
func isPowerOfTwo(n int) bool {
	const big = 1 << 30
	return n > 0 && big%n == 0
}
