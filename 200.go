package partice

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var result int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				affect(grid, i, j)
				result++
			}
		}
	}
	return result
}

func affect(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}
	if grid[x][y] == '1' {
		grid[x][y] = '2'
		affect(grid, x-1, y)
		affect(grid, x, y-1)
		affect(grid, x+1, y)
		affect(grid, x, y+1)
	}
}
