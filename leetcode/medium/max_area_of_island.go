package main

import "fmt"

func main() {
	fmt.Println()
	a := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(a))
}

/*
@link https://leetcode.cn/problems/max-area-of-island/
岛屿的最大面积
1. 遍历整个二维数组，当值等于1时进行处理
2. 对等于1的元素上下左右遍历
3. 对多个count值进行取最大值
*/
func maxAreaOfIsland(grid [][]int) int {
	maxV := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxV = maxVal(maxV, helper(grid, i, j))
			}
		}
	}
	return maxV
}
func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 处理过的元素将其赋值为0 避免二次遍历
func helper(grid [][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	count := 1
	// 处理左边
	count += helper(grid, i-1, j)
	// 处理右边
	count += helper(grid, i+1, j)
	// 处理上边
	count += helper(grid, i, j-1)
	// 处理下边
	count += helper(grid, i, j+1)
	return count
}
