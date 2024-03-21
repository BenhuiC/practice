package array

func surfaceArea(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dir := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			h := grid[i][j]
			if h == 0 {
				continue
			}
			area := h*4 + 2
			for _, d := range dir {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < m && y >= 0 && y < n {
					area -= min(h, grid[x][y])
				}
			}
			res += area
		}
	}
	return res
}
