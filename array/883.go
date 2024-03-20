package array

func projectionArea(grid [][]int) int {
	res := 0
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	for i := 0; i < m; i++ {
		left := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res++
			left = max(left, grid[i][j])
		}
		res += left
	}

	for i := 0; i < n; i++ {
		front := 0
		for j := 0; j < m; j++ {
			front = max(front, grid[j][i])
		}
		res += front
	}

	return res
}
