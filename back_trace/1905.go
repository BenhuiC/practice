package back_trace

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	if m == 0 {
		return 0
	}
	n := len(grid1[0])
	direction := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	res := 0
	var bfs func(x, y int)
	bfs = func(x, y int) {
		ary := make([][2]int, 0)
		ary = append(ary, [2]int{x, y})
		grid2[x][y] = -1
		fg := false
		for len(ary) != 0 {
			next := make([][2]int, 0)
			for _, a := range ary {
				for _, d := range direction {
					x1, y1 := a[0]+d[0], a[1]+d[1]
					if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
						continue
					}
					if grid2[x1][y1] == 1 {
						if grid1[x1][y1] == 0 {
							fg = true
						}
						grid2[x1][y1] = -1
						next = append(next, [2]int{x1, y1})
					}
				}
			}
			ary = next
		}
		if !fg {
			res++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	return res
}
