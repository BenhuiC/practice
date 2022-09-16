package partice

func sumOfDistancesInTree(n int, edges [][]int) []int {
	// todo
	res := make([]int, n)
	if n == 1 {
		return res
	}
	if n == 2 {
		return []int{1, 1}
	}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, i+1)
	}
	for _, v := range edges {
		if v[0] > v[1] {
			matrix[v[0]][v[1]] = 1
		} else {
			matrix[v[1]][v[0]] = 1
		}
	}
	var setLen func(x, y, v int)
	setLen = func(x, y, v int) {
		if x < y {
			x, y = y, x
		}
		matrix[x][y] = v
	}
	var nodeLen func(oriNode, x, y int) int
	nodeLen = func(oriNode, x, y int) int {
		if x == y {
			return 0
		}
		if x > y && matrix[x][y] != 0 {
			return matrix[x][y]
		} else if x < y && matrix[y][x] != 0 {
			return matrix[y][x]
		}

		for i, v := range matrix[x] {
			if v == 0 || i == oriNode {
				continue
			}
			if t := nodeLen(oriNode, i, y); t != 0 {
				setLen(x, y, t+1)
				return t + 1
			}
		}
		for idx := x; idx < n; idx++ {
			v := matrix[idx][x]
			if v == 0 || idx == oriNode {
				continue
			}
			if t := nodeLen(oriNode, idx, y); t != 0 {
				setLen(x, y, t+1)
				return t + 1
			}
		}
		return 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i] += nodeLen(i, i, j)
		}
	}
	printMatrix(matrix)

	return res
}
