package _59_spiral_matrix_ii

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i, _ := range ret {
		ret[i] = make([]int, n)
	}
	// 左上
	top := 0
	left := 0

	// 右下
	right := n - 1
	bottom := n - 1

	// 初始值
	num := 1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ret[top][i] = num
			num++
		}
		for i := top + 1; i <= bottom; i++ {
			ret[i][right] = num
			num++
		}
		for i := right - 1; i >= left; i-- {
			ret[bottom][i] = num
			num++
		}
		for i := bottom - 1; i > top; i-- {
			ret[i][left] = num
			num++
		}
		top++
		left++
		right--
		bottom--
	}
	return ret
}
