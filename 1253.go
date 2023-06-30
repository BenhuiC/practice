package partice

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res1253 := make([][]int, 2)
	l1 := make([]int, len(colsum))
	l2 := make([]int, len(colsum))
	colsums := 0
	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 2 {
			upper--
			lower--
			l1[i] = 1
			l2[i] = 1
		} else {
			colsums += colsum[i]
		}
	}
	if colsums != upper+lower || upper < 0 || lower < 0 {
		return [][]int{}
	}
	for i := 0; i < len(colsum); i++ {
		if colsum[i] != 1 {
			continue
		}
		if upper > 0 {
			upper--
			l1[i] = 1
		} else if lower > 0 {
			lower--
			l2[i] = 1
		} else {
			return [][]int{}
		}
	}
	if upper > 0 || lower > 0 {
		return [][]int{}
	}
	res1253[0], res1253[1] = l1, l2

	return res1253
}
