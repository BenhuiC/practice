package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		l, h := 0, len(matrix[0])
		tmp := matrix[i]
		if tmp[0] > target || tmp[h-1] < target {
			continue
		}
		for l < h {
			m := (l + h) / 2
			if tmp[m] == target {
				return true
			} else if tmp[m] > target {
				h = m
			} else {
				l = m + 1
			}
		}
	}
	return false
}

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := 0, len(matrix[0])-1
	var v int
	for m < len(matrix) && n >= 0 {
		v = matrix[m][n]
		if v == target {
			return true
		} else if v < target {
			m++
		} else {
			n--
		}
	}
	return false
}
