package offer

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	if m == 0 {
		return nil
	}
	n := len(img[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	avg := func(x, y int) int {
		directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		num, sum := 0, 0
		for _, d := range directions {
			tx, ty := x+d[0], y+d[1]
			if tx >= 0 && tx < m && ty >= 0 && ty < n {
				num++
				sum += img[tx][ty]
			}
		}
		return sum / num
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = avg(i, j)
		}
	}
	return res
}
