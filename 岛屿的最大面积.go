package pos

func maxAreaOfIsland(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	area := 0
	row, col := len(grid), len(grid[0])

	// 水平或竖直方向
	direction := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				area = max(area, dfs_land(i, j, row, col, &grid, direction))
			}
		}
	}
	return area
}

func dfs_land(curX, curY, row, col int, grid *[][]int, direction [][]int) int {
	// if curX<0 || curX>=row || curY<0 || curY>=col || (*grid)[curX][curY]==0 {
	// 	return 0
	// }
	(*grid)[curX][curY] = 0
	ans := 1
	for i := 0; i < 4; i++ {
		newX := curX + direction[i][0]
		newY := curY + direction[i][1]
		ans += dfs_land(newX, newY, row, col, grid, direction)
	}
	return ans
}
