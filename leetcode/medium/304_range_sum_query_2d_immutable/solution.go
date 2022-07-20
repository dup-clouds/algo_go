package _304_range_sum_query_2d_immutable

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m < 1 {
		return NumMatrix{}
	}
	sums := make([][]int, m+1)
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{sums: sums}
	}
	sums[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		sums[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			sums[i][j] = matrix[i-1][j-1] + sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1]
		}
	}
	return NumMatrix{sums: sums}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.sums[row2+1][col2+1] - nm.sums[row1][col2+1] - nm.sums[row2+1][col1] + nm.sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
