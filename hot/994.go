package hot

func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])

	maxStep := 2
	var mark func(i, j, step int)
	mark = func(i, j, step int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || (grid[i][j] != 1 && grid[i][j] < step) {
			return
		}
		grid[i][j] = step
		step++
		mark(i-1, j, step)
		mark(i+1, j, step)
		mark(i, j-1, step)
		mark(i, j+1, step)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				mark(i, j, 2)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			} else if grid[i][j] == 1 {
				return -1
			} else if grid[i][j] > maxStep {
				maxStep = grid[i][j]
			}
		}
	}
	return maxStep - 2
}
