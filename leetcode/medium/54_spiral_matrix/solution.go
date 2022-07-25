package _54_rotate_image

func spiralOrder(matrix [][]int) []int {
	rowLen := len(matrix)
	columnLen := len(matrix[0])
	var ret []int
	// 左上角 负责从左到右
	top := 0
	// 右上角 负责从上到下
	right := columnLen - 1
	// 右下角 负责从右到左
	bottom := rowLen - 1
	// 左下角 负责从下到上
	left := 0
	// 确保所有元素遍历完成
	for len(ret) < rowLen*columnLen {
		//  从左到右 确保不超出行的限制，即保证遍历在第一行到最后一行
		if top <= bottom {
			// 从左到右遍历行，结束点为右边界
			for i := left; i <= right; i++ {
				ret = append(ret, matrix[top][i])
			}
			// 上方行+1
			top++
		}
		//  从右到下 确保不超出列的限制
		if left <= right {
			// 从上到下遍历列
			for j := top; j <= bottom; j++ {
				ret = append(ret, matrix[j][right])
			}
			// 右边界-1
			right--
		}

		//	从右到左,确保不超出行的限制
		if top <= bottom {
			for i := right; i >= left; i-- {
				ret = append(ret, matrix[bottom][i])
			}
			// 下边界-1
			bottom--
		}
		//	从下到上，确保不超出列限制
		if left <= right {
			for j := bottom; j >= top; j-- {
				ret = append(ret, matrix[j][left])
			}
			// 左边界+1
			left++
		}
	}
	return ret
}

func spiralOrder2(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ret := make([]int, m*n)
	// ret数组索引维护
	index := 0
	// 左上
	top := 0
	left := 0
	// 右下
	right := n - 1
	bottom := m - 1

	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			ret[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			ret[index] = matrix[row][right]
			index++
		}
		// 防止单行层级遍历越界
		if left < right && top < bottom {
			for column := right - 1; column >= left; column-- {
				ret[index] = matrix[bottom][column]
				index++
			}
			for row := bottom - 1; row > top; row-- {
				ret[index] = matrix[row][left]
				index++
			}
		}
		left++
		top++
		bottom--
		right--
	}
	return ret
}
