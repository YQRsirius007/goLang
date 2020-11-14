package pos

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	rows, cols := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, rows-1, 0, cols-1
	total := rows * cols

	for i := 0; i < total; {

		for col := left; i < total && col <= right; col++ {
			res = append(res, matrix[top][col])
			i++
		}
		top++

		for row := top; i < total && row <= bottom; row++ {
			res = append(res, matrix[row][right])
			i++
		}
		right--

		for col := right; i < total && col >= left; col-- {
			res = append(res, matrix[bottom][col])
			i++
		}
		bottom--

		for row := bottom; i < total && row >= top; row-- {
			res = append(res, matrix[row][left])
			i++
		}
		left++

	}
	return res
}
