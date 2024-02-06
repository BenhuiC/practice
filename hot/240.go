package hot

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	findInAry := func(ary []int) bool {
		l, r := 0, len(ary)-1
		for l <= r {
			mid := (r-l)/2 + l
			if ary[mid] < target {
				l = mid + 1
			} else if ary[mid] > target {
				r = mid - 1
			} else {
				return true
			}
		}

		return false
	}

	for i := 0; i < m; i++ {
		if matrix[i][0] > target {
			break
		}
		if matrix[i][n-1] < target {
			continue
		}
		if findInAry(matrix[i]) {
			return true
		}
	}

	return false
}
