package offer

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	var calDis func(i1, i2, j1, j2 int) int
	calDis = func(i1, j1, i2, j2 int) int {
		if i1 > j2 || j1 < i2 {
			return 0
		}

		i, j := i1, j1
		if i1 < i2 {
			i = i2
		}
		if j2 < j1 {
			j = j2
		}

		return j - i
	}

	sameArea := calDis(ax1, ax2, bx1, bx2) * calDis(ay1, ay2, by1, by2)

	res := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - sameArea

	return res
}
