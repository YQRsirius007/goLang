package pos

import "fmt"

func gameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])
	direction := []int{0, 1, -1}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			living := 0

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !(direction[i] == 0 && direction[j] == 0) {
						r := row + direction[i]
						c := col + direction[j]
						if (r < rows && r >= 0) && (c < cols && c >= 0) && (board[r][c] == 1 || board[r][c] == -1) {
							living++
						}
					}
				}
			}
			fmt.Println(row, col, living)
			if board[row][col] == 1 && (living < 2 || living > 3) {
				board[row][col] = -1
			}
			if board[row][col] == 0 && living == 3 {
				board[row][col] = 2
			}
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] > 0 {
				board[row][col] = 1
			} else {
				board[row][col] = 0
			}
		}
	}
}
