package partice

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	low, high := 0, m
	var mp []int
	for low < high {
		t := (low + high) / 2
		if matrix[t][0] <= target && matrix[t][n-1] >= target {
			mp = matrix[t]
			break
		} else if matrix[t][0] > target {
			high = t
		} else {
			low = t + 1
		}
	}
	if mp == nil || len(mp) == 0 {
		return false
	}
	l, h := 0, n
	for l < h {
		tp := (l + h) / 2
		if mp[tp] == target {
			return true
		} else if mp[tp] > target {
			h = tp
		} else {
			l = tp + 1
		}
	}

	return false
}
