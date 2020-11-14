package pos

func slove(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])
	direction := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && board[i][j] == 'O' {
				dfs_slove(&board, i, j, rows, cols, direction)
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
	return
}
func dfs_slove(board *[][]byte, x, y, rows, cols int, direction [][]int) {
	if x < 0 || x >= rows || y < 0 || y >= cols || (*board)[x][y] == 'X' || (*board)[x][y] == '#' {
		return
	}
	(*board)[x][y] = '#'
	for i := 0; i < 4; i++ {
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		dfs_slove(board, newX, newY, rows, cols, direction)
	}
}
