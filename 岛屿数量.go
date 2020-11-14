package pos

// dfs深度优先遍历 把连接的岛屿全部标记一遍
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	landNumber := 0
	marked := make([][]int, row)
	// direction := make([][]int, 4)
	// direction[0] = []int{-1, 0}
	// direction[1] = []int{0, 1}
	// direction[2] = []int{1, 0}
	// direction[3] = []int{0, -1}
	direction := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for i := 0; i < row; i++ {
		marked[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' && marked[i][j] == 0 {
				landNumber++
				dfs(i, j, row, col, grid, &marked, direction)
			}
		}
	}
	return landNumber
}
func dfs(curX, curY, row, col int, grid [][]byte, marked *[][]int, direction [][]int) {
	(*marked)[curX][curY] = 1
	for i := 0; i < 4; i++ {
		newX := curX + direction[i][0]
		newY := curY + direction[i][1]
		if (newX >= 0 && newX < row && newY >= 0 && newY < col) && grid[newX][newY] == '1' && (*marked)[newX][newY] == 0 {
			dfs(newX, newY, row, col, grid, marked, direction)
		}
	}
}
