package back_trace

import "fmt"

func numEnclaves(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])
	if m == 1 || n == 1 {
		return 0
	}
	direction := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var bfs func(x, y int)
	bfs = func(x, y int) {
		nodeNu := 0
		ary := make([][2]int, 0)
		ary = append(ary, [2]int{x, y})
		grid[x][y] = -1

		fg := false
		for len(ary) != 0 {
			nodeNu += len(ary)
			next := make([][2]int, 0)
			for _, v := range ary {
				for _, d := range direction {
					x1, y1 := v[0]+d[0], v[1]+d[1]
					if x1 < 0 || x1 >= m || y1 < 0 || y1 >= n {
						fg = true
					} else if grid[x1][y1] == 1 {
						grid[x1][y1] = -1
						next = append(next, [2]int{x1, y1})
					}
				}
			}
			ary = next
		}
		if !fg {
			res += nodeNu
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
				printMatrix(grid)
				fmt.Println(res)
			}
		}
	}

	return res
}
