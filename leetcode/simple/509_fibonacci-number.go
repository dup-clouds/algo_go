package main

import "fmt"

func main() {
	fmt.Println(fib1(10))
}

/**
@link https://leetcode.cn/problems/fibonacci-number/
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
*/
func fib(n int) int {
	mem := make(map[int]int)
	mem[0] = 0
	mem[1] = 1
	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}

func fib1(n int) int {
	mem := make(map[int]int)
	return fibRec(n, &mem)
}
func fibRec(n int, mem *map[int]int) int {
	if (*mem)[n] > 0 {
		return (*mem)[n]
	}
	if n < 2 {
		return n
	}
	res := fibRec(n-1, mem) + fibRec(n-2, mem)
	(*mem)[n] = res
	return res
}

func fib3(n int) int {
	if n < 2 {
		return n
	}
	prev := 0
	curr := 1
	for i := 2; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
