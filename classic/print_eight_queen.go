package main

import "fmt"

func main() {
	ans := [8]int{}
	count := 0
	cal8queens(0, &ans, &count)
	fmt.Println(count)
}

/**
8皇后问题所有可行方案
@link https://time.geekbang.org/column/article/74287
*/
func cal8queens(row int, ans *[8]int, count *int) {
	if row == 8 {
		printQueen(*ans)
		fmt.Println()
		fmt.Println(ans)
		fmt.Println()
		*count++
		return
	}
	for column := 0; column < 8; column++ {
		if isValid(*ans, row, column) {
			ans[row] = column
			cal8queens(row+1, ans, count)
		}
	}
}

func isValid(ans [8]int, row int, column int) bool {
	lowColumn := column - 1
	highColumn := column + 1
	// 判断当前行与当前已选皇后位置是否存在冲突
	for i := row - 1; i >= 0; i-- {
		// 正上方
		if ans[i] == column {
			return false
		}
		// 左上
		if lowColumn >= 0 {
			if ans[i] == lowColumn {
				return false
			}
		}
		// 右上
		if highColumn < 8 {
			if ans[i] == highColumn {
				return false
			}
		}
		// 列随着row递减或递增
		lowColumn--
		highColumn++
	}
	return true
}

func printQueen(ans [8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if ans[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
