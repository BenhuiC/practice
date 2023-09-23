package back_trace

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	direction := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	res := 0
	var bfs func(x, y int)
	bfs = func(x, y int) {
		ary := make([][2]int, 0)
		ary = append(ary, [2]int{x, y})
		grid[x][y] = 'x'
		for len(ary) != 0 {
			next := make([][2]int, 0)
			for _, a := range ary {
				for _, d := range direction {
					x1, y1 := a[0]+d[0], a[1]+d[1]
					if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
						continue
					}
					if grid[x1][y1] == '1' {
						next = append(next, [2]int{x1, y1})
						grid[x1][y1] = 'x'
					}
				}
			}
			ary = next
		}
		res++
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				bfs(i, j)
			}
		}
	}

	return res
}
