package offer

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	var i, j int
	for _, v := range ops {
		if i == 0 || v[0] < i {
			i = v[0]
		}
		if j == 0 || v[1] < j {
			j = v[1]
		}

	}

	return i * j
}
