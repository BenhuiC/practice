package offer

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if k > n*n {
		return -1
	}
	ary := make([][2]int, 0, n)
	ma, na := make(map[int]int, n), make(map[int]int, n)
	nu := 0
	ary = append(ary, [2]int{0, 0})
	for len(ary) != 0 {
		idx := 0
		for i := 1; i < len(ary); i++ {
			if matrix[ary[i][0]][ary[i][1]] < matrix[ary[idx][0]][ary[idx][1]] {
				idx = i
			}
		}
		x, y := ary[idx][0], ary[idx][1]
		ary[idx] = ary[len(ary)-1]
		ary = ary[:len(ary)-1]
		nu++
		if nu == k {
			return matrix[x][y]
		}
		ma[x] = y
		na[y] = x
		if y+1 < n {
			if x == 0 {
				ary = append(ary, [2]int{x, y + 1})
			} else if v, ok := na[y+1]; ok && v == x-1 {
				ary = append(ary, [2]int{x, y + 1})
			}
		}
		if x+1 < n {
			if y == 0 {
				ary = append(ary, [2]int{x + 1, y})
			} else if v, ok := ma[x+1]; ok && v == y-1 {
				ary = append(ary, [2]int{x + 1, y})
			}
		}
	}
	return 0
}
