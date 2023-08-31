package partice

func minPathSum(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	return minSum(grid, 0, 0)
}

func minSum(grid [][]int, x, y int) int {
	if x >= len(grid) || y >= len(grid[0]) {
		return -1
	}
	l, h := minSum(grid, x+1, y), minSum(grid, x, y+1)

	if l < 0 && h < 0 {
		return grid[x][y]
	}
	if l < 0 && h >= 0 {
		return h + grid[x][y]
	}
	if h < 0 && l >= 0 {
		return l + grid[x][y]
	}

	return Min(l, h) + grid[x][y]
}

func minPathSum2(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	mp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		mp[i] = make([]int, len(grid[i]))
	}
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if x-1 < 0 && y-1 < 0 {
				mp[x][y] = grid[x][y]
			} else if x-1 < 0 {
				mp[x][y] = grid[x][y] + mp[x][y-1]
			} else if y-1 < 0 {
				mp[x][y] = grid[x][y] + mp[x-1][y]
			} else {
				mp[x][y] = Min(mp[x-1][y], mp[x][y-1]) + grid[x][y]
			}
		}
	}
	return mp[len(grid)-1][len(grid[0])-1]
}
