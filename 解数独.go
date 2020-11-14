package pos

func solveSudoku(board [][]byte) {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	dfs_sudo(&board, 0, 0)
}
func dfs_sudo(board *[][]byte, row int, col int) bool {
	if col == 9 {
		return dfs_sudo(board, row+1, 0)
	}
	if row == 9 {
		return true
	}
	if (*board)[row][col] != '.' {
		return dfs_sudo(board, row, col+1)
	}
	var c byte
	for c = '1'; c <= '9'; c++ {
		if !isValid(*board, row, col, c) {
			continue
		}
		(*board)[row][col] = byte(c)
		if dfs_sudo(board, row, col+1) {
			return true
		}
		(*board)[row][col] = '.'
	}
	return false
}
func isValid(board [][]byte, row int, col int, c byte) bool {
	for k := 0; k < 9; k++ {
		if board[row][k] == c {
			return false
		}
		if board[k][col] == c {
			return false
		}
		if board[(row/3)*3+k/3][(col/3)*3+k%3] == c {
			return false
		}
	}
	return true
}
