package pos

// 找出公式后，用水平翻转，对角线反转，置换等方式分步实现
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
