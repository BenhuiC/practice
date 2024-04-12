package array

func shortestBridge(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	type point struct {
		x, y int
	}
	direct := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	findBord := func(x, y int) []point {
		ary := make([]point, 0)
		ary = append(ary, point{x, y})
		grid[x][y] = 2
		res := make([]point, 0)

		for len(ary) != 0 {
			next := make([]point, 0)
			for _, a := range ary {
				add := false
				for _, d := range direct {
					i, j := a.x+d[0], a.y+d[1]
					if i < 0 || i >= m || j < 0 || j >= n {
						continue
					}
					if grid[i][j] == 1 {
						grid[i][j] = 2
						next = append(next, point{i, j})
					} else if grid[i][j] == 0 && !add {
						add = true
						res = append(res, a)
					}
				}
			}
			ary = next
		}
		return res
	}

	findDis := func(board []point) int {
		res := 0

		for len(board) != 0 {
			next := make([]point, 0)
			for _, b := range board {
				for _, d := range direct {
					i, j := b.x+d[0], b.y+d[1]
					if i < 0 || i >= m || j < 0 || j >= n {
						continue
					}

					if grid[i][j] == 0 {
						grid[i][j] = 2
						next = append(next, point{i, j})
					} else if grid[i][j] == 1 {
						return res
					}
				}
			}

			res++
			board = next
		}

		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				board := findBord(i, j)
				return findDis(board)
			}
		}
	}

	return 0
}
