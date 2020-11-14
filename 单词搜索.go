package pos

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	rows, cols := len(board), len(board[0])
	if rows == 0 || cols == 0 {
		return false
	}
	direction := [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
	visited := make([][]int, rows)
	for i := range visited {
		visited[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			path := string(word[0])
			if dfs_exist(board, &visited, direction, i, j, rows, cols, word, &path) {
				return true
			}
		}
	}
	return false
}
func dfs_exist(board [][]byte, visited *[][]int, direction [][]int, x, y, rows, cols int, word string, path *string) bool {
	if board[x][y] != word[len(*path)-1] {
		return false
	}
	if len(word) == len(*path) {
		return true
	}
	(*visited)[x][y] = 1
	for i := 0; i < 4; i++ {
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && (*visited)[newX][newY] == 0 {
			*path += string(word[len(*path)])
			if dfs_exist(board, visited, direction, newX, newY, rows, cols, word, path) {
				return true
			}
			*path = (*path)[:len(*path)-1]
		}
	}
	(*visited)[x][y] = 0
	return false
}
