package offer

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
	}
	zeroesPos := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroesPos = append(zeroesPos, []int{i, j})
			}
		}
	}
	q := [][]int{}
	q = append(q, zeroesPos...)
	seen := make(map[[2]int]bool)
	for _, pos := range zeroesPos {
		seen[[2]int{pos[0], pos[1]}] = true
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		i, j := current[0], current[1]
		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !seen[[2]int{ni, nj}] {
				dist[ni][nj] = dist[i][j] + 1
				q = append(q, []int{ni, nj})
				seen[[2]int{ni, nj}] = true
			}
		}
	}

	return dist
}
