package offer

func islandPerimeter(grid [][]int) int {
	res := 0
	dir := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for _, v := range dir {
				x, y := i+v[0], j+v[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					res++
					continue
				}
				if grid[x][y] == 0 {
					res++
				}
			}
		}
	}

	return res
}
