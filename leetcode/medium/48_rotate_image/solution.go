package _48_rotate_image

// 原始方法
func rotate(matrix [][]int) {
	temp := make([][]int, len(matrix))
	for i, _ := range temp {
		t := make([]int, len(matrix[0]))
		temp[i] = t
		copy(t, matrix[i])
	}
	for i := 0; i < len(matrix); i++ {
		rowTemp := make([]int, 0)
		for j := len(matrix) - 1; j >= 0; j-- {
			rowTemp = append(rowTemp, temp[j][i])
		}
		matrix[i] = rowTemp
	}
}

// 优化rotate方法
func rotate1(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i, _ := range temp {
		temp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			// 原有的第i行第j列 旋转为现有的第j行倒数第i列
			temp[j][n-1-i] = v
		}
	}
	copy(matrix, temp)
}

// 先主对角线对折，再按行反转
func rotate2(matrix [][]int) {
	// 主对角线对折
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 按行反转
	for i := 0; i < len(matrix); i++ {
		reverse(matrix[i])
	}
}

func reverse(row []int) {
	left := 0
	right := len(row) - 1
	for left < right {
		row[left], row[right] = row[right], row[left]
		left++
		right--
	}
}
