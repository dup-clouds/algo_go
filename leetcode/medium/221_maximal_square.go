package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(matrix))
}

// @link https://leetcode.cn/problems/maximal-square/
//221. 最大正方形
func maximalSquare(matrix [][]byte) int {
	max := 0
	rowLen := len(matrix)
	columnLen := len(matrix[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			if matrix[i][j] == '1' {
				max = maxSquare(max, 1)
				maxCycle := minSquare(rowLen-i, columnLen-j)
				for k := 1; k < maxCycle; k++ {
					if matrix[i+k][j+k] != '1' {
						break
					}
					flag := true
					for g := 0; g < k; g++ {
						if matrix[i+k][j+g] == '0' || matrix[i+g][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						max = maxSquare(k+1, max)
					} else {
						break
					}
				}
			}
		}
	}
	return max * max
}

func maxSquare(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func minSquare(x, y int) int {
	if x < y {
		return x
	}
	return y
}
