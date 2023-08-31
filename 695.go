package partice

func maxAreaOfIsland(grid [][]int) (res int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			a := area(grid, i, j)
			if a > res {
				res = a
			}
		}
	}
	//printMatrix(grid)
	return
}

func area(grid [][]int, i, j int) (a int) {
	t := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m := len(grid)
	n := len(grid[0])
	q := [][2]int{}
	q = append(q, [2]int{i, j})
	grid[i][j] = -1
	for {
		if a >= len(q) {
			break
		}
		var curX, curY int
		for _, v := range t {
			curX, curY = q[a][0]+v[0], q[a][1]+v[1]
			if curX < 0 || curX >= m || curY < 0 || curY >= n || grid[curX][curY] != 1 {
				continue
			}
			grid[curX][curY] = -1
			q = append(q, [2]int{curX, curY})
		}
		a++
	}
	return
}
