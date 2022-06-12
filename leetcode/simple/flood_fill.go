package main

import "fmt"

func main() {
	fmt.Println()
	a := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(a, 1, 1, 2))
}

/**
@link https://leetcode.cn/problems/flood-fill/
733 图像渲染
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	helper(image, sr, sc, newColor, image[sr][sc])
	return image
}

/**
1. 新颜色和旧颜色一致时则无需染色
2. sr，sc其中一个出边界时停止染色
3. 从上下左右进行染色
*/
func helper(image [][]int, sr int, sc int, newColor int, oldColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != oldColor || newColor == oldColor {
		return
	}
	image[sr][sc] = newColor
	// 处理左边
	helper(image, sr-1, sc, newColor, oldColor)
	// 处理右边
	helper(image, sr+1, sc, newColor, oldColor)
	// 处理上边
	helper(image, sr, sc-1, newColor, oldColor)
	// 处理下边
	helper(image, sr, sc+1, newColor, oldColor)
}
