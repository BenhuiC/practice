package partice

func maxIncreaseKeepingSkyline(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	res807 := 0
	lheight := make([]int, m)
	bheight := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > lheight[i] {
				lheight[i] = grid[i][j]
			}
			if grid[i][j] > bheight[j] {
				bheight[j] = grid[i][j]
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			h := min(lheight[i], bheight[j])
			res807 += h - grid[i][j]
		}
	}

	return res807
}
