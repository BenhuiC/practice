package offer

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	sColor := image[sr][sc]
	if sColor == color {
		return image
	}
	m, n := len(image), len(image[0])
	dir := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	ary := [][2]int{{sr, sc}}
	for len(ary) != 0 {
		q := ary[0]
		ary = ary[1:]
		image[q[0]][q[1]] = color
		for _, d := range dir {
			x, y := q[0]+d[0], q[1]+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && image[x][y] == sColor {
				ary = append(ary, [2]int{x, y})
			}
		}
	}
	return image
}
