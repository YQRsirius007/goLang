package pos

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix)+1)
	maxL := 0
	leftL, topL, leftAndTopL, curL := 0, 0, 0, 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				curL = min(min(leftL, topL), leftAndTopL) + 1
				if curL > maxL {
					maxL = curL
				}
			}

		}
	}
	return maxL * maxL
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
